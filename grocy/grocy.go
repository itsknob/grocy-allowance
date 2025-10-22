package grocy

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type GrocyConfig struct {
	// GROCY_API_KEY string
	GROCY_URL string
}

type GrocyApi interface {
	GetStock() ([]StockEntry, error)
}

type GrocyClient struct {
	baseUrl string
	client  *http.Client
}

func (c *GrocyClient) InitAllowance() {

	body := []byte(`{ "json": "body"}`)

	// Set up Request to create new User Entity
	req, err := http.NewRequest("POST", "/objects/userentities", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Failed to create new request for new user entity")
	}

	// Add Headers
	req.Header.Add("GROCY-API-KEY", os.Getenv("GROCY_API_KEY"))
	req.Header.Add("accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal("Failed to create new user entity: Allowance")
	}
	if resp.StatusCode == 200 {
		return
	} else {
		log.Fatal("Request to create new user entity failed.")
	}
	log.Default().Print("Created 'Allowance' user entity")
}

func (c *GrocyClient) HasAllowance() bool {
	req, err := c.createRequestWithHeaders("GET", "/objects/userentities?query[]=name=Allowance", nil)
	if err != nil {
		return false
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	type UserEntitiesQueryResponse struct {
		Id                int    `json:"id"`
		Name              string `json:"name"`
		Caption           string `json:"caption"`
		Description       string `json:"description"`
		ShowInSidebarMenu int    `json:"show_in_sidebar_menu"`
		IconCssClass      string `json:"icon_css_class"`
		CreatedDate       string `json:"row_created_timestamp"`
	}

	var temp []UserEntitiesQueryResponse
	if err := json.NewDecoder(resp.Body).Decode(&temp); err != nil {
		return false
	}

	if len(temp) > 0 {
		return true
	} else {
		return false
	}
}

func (c *GrocyClient) GetUnits() string {
	c.createRequestWithHeaders("GET", "/", nil)
	panic("unimplemented")
}

func (c *GrocyClient) createRequestWithHeaders(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseUrl+path, body)
	if err != nil {
		return nil, err
	}

	// Add Headers
	req.Header.Add("GROCY-API-KEY", os.Getenv("GROCY_API_KEY"))
	req.Header.Add("accept", "application/json")

	return req, nil
}

func NewGrocyClient(baseUrl string) *GrocyClient {
	config := GrocyConfig{
		os.Getenv("GROCY_URL"),
	}

	c := &GrocyClient{
		baseUrl: config.GROCY_URL,
		client:  &http.Client{},
	}

	if c.HasAllowance() {
		return c
	} else {
		c.InitAllowance()
		return c
	}
}
