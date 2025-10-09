package grocy

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Product struct {
	Id                                      int
	Name                                    string
	Description                             string
	Location_id                             int
	Qu_id_purchase                          int
	Qu_id_stock                             int
	Enable_tare_weight_handling             int
	Not_check_stock_fulfillment_for_recipes int
	Product_group_id                        int
	Tare_weight                             float32
	Min_stack_amount                        uint // default 0
	Default_best_before_days                uint // default 0
	Default_best_before_days_after_open     uint //default 0
	Picture_file_name                       string
	Row_created_timestamp                   string // datetime
	Shopping_location_id                    int
	Treat_opened_as_out_of_stock            int
	Auto_reprint_stock_label                int
	No_own_stock                            int
	Should_not_be_frozen                    int
	Default_consumer_location_id            int
	Move_on_open                            int
}

type Item struct {
	Product_id               int
	Amount                   int8
	Amount_aggregated        int8
	Amount_opened            int8
	Amount_opened_aggregated int8
	Best_before_date         string // datetime
	Is_aggregated_amount     int
	Product                  Product
}

type GrocyConfig struct {
	// GROCY_API_KEY string
	GROCY_URL string
}

type GrocyApi interface {
	GetStock() ([]Item, error)
}

type GrocyClient struct {
	baseUrl string
	client  *http.Client
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

	return &GrocyClient{
		baseUrl: config.GROCY_URL,
		client:  &http.Client{},
	}
}

func (c *GrocyClient) GetStock() (*[]Item, error) {
	// Set up request
	req, err := c.createRequestWithHeaders("GET", "/stock", nil)
	if err != nil {
		return nil, err
	}

	// Fire Request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Parse Response
	var stock []Item
	if err := json.NewDecoder(resp.Body).Decode(&stock); err != nil {
		return nil, err
	}

	return &stock, nil
}
