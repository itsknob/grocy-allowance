package grocy

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// type Product struct {
// 	Id                                      int     `json:"id"`
// 	Name                                    string  `json:"name"`
// 	Description                             string  `json:"description"`
// 	Location_id                             int     `json:"location_id"`
// 	Qu_id_purchase                          int     `json:"qu_id_purchase"`
// 	Qu_id_stock                             int     `json:"qu_id_stock"`
// 	Enable_tare_weight_handling             int     `json:"enable_tare_weight_handling"`
// 	Not_check_stock_fulfillment_for_recipes int     `json:"not_check_stock_fulfillment_for_recipes"`
// 	Product_group_id                        int     `json:"product_group_id"`
// 	Tare_weight                             float32 `json:"tare_weight"`
// 	Min_stack_amount                        uint    `json:"min_stack_amount"`
// 	Default_best_before_days                uint    `json:"default_best_before_days"`
// 	Default_best_before_days_after_open     uint    `json:"default_best_before_days_after_open"`
// 	Picture_file_name                       string  `json:"picture_file_name"`
// 	Row_created_timestamp                   string  `json:"row_created_timestamp"`
// 	Shopping_location_id                    int     `json:"shopping_location_id"`
// 	Treat_opened_as_out_of_stock            int     `json:"treat_opened_as_out_of_stock"`
// 	Auto_reprint_stock_label                int     `json:"auto_reprint_stock_label"`
// 	No_own_stock                            int     `json:"no_own_stock"`
// 	Should_not_be_frozen                    int     `json:"should_not_be_frozen"`
// 	Default_consumer_location_id            int     `json:"default_consumer_location_id"`
// 	Move_on_open                            int     `json:"move_on_open"`
// }
//
// type Item struct {
// 	Product_id               int `json:"product_id"`
// 	Amount                   int `json:"amount"`
// 	Amount_aggregated        int `json:"amount_aggregated"`
// 	Amount_opened            int `json:"amount_opened"`
// 	Amount_opened_aggregated int `json:"amount_opened_aggregated"`
// 	Best_before_date         string `json:"best_before_date"`
// 	Is_aggregated_amount     int `json:"is_aggregated_amount"`
// 	Product                  Product `json:"product"`
// }
//
// type CustomUserEntityRequest struct {
// 	Name                 string
// 	Caption              string
// 	Description          string
// 	Show_in_sidebar_menu string
// 	Icon_css_class       string
// }

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
		ShowInSidebarMenu int    `json:"Show_in_sidebar_menu"`
		IconCssClass      string `json:"icon_css_class"`
		CreatedDate       string `json:"row_created_timestamp"`
	}

	var temp []byte
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

	return &GrocyClient{
		baseUrl: config.GROCY_URL,
		client:  &http.Client{},
	}
}

func (c *GrocyClient) GetStock() (*[]StockEntry, error) {
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
	var stock []StockEntry
	if err := json.NewDecoder(resp.Body).Decode(&stock); err != nil {
		return nil, err
	}

	return &stock, nil
}

func (c *GrocyClient) createCustomUserEntity() {
	log.Fatal("Not implemented")
}
