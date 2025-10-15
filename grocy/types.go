package grocy

import (
	"time"
)

// Generated from openapi.json components.schemas

type StockTransactionType int

const (
	PURCHASE StockTransactionType = iota
	CONSUME
	INVENTORY_CORRECTION
	PRODUCT_OPENED
)

var TransactionType = map[StockTransactionType]string{
	PURCHASE:             "purchase",
	CONSUME:              "consume",
	INVENTORY_CORRECTION: "inventory-correction",
	PRODUCT_OPENED:       "product-opened",
}

type Product struct {
	Id                                 int            `json:"id,omitempty"`
	Name                               string         `json:"name,omitempty"`
	Description                        string         `json:"description,omitempty"`
	LocationId                         int            `json:"location_id,omitempty"`
	QuIdPurchase                       int            `json:"qu_id_purchase,omitempty"`
	QuIdStock                          int            `json:"qu_id_stock,omitempty"`
	EnableTareWeightHandling           int            `json:"enable_tare_weight_handling,omitempty"`
	NotCheckStockFulfillmentForRecipes int            `json:"not_check_stock_fulfillment_for_recipes,omitempty"`
	ProductGroupId                     int            `json:"product_group_id,omitempty"`
	TareWeight                         float64        `json:"tare_weight,omitempty"`
	MinStockAmount                     float64        `json:"min_stock_amount,omitempty"`
	DefaultBestBeforeDays              int            `json:"default_best_before_days,omitempty"`
	DefaultBestBeforeDaysAfterOpen     int            `json:"default_best_before_days_after_open,omitempty"`
	PictureFileName                    string         `json:"picture_file_name,omitempty"`
	RowCreatedTimestamp                time.Time      `json:"row_created_timestamp"`
	ShoppingLocationId                 int            `json:"shopping_location_id,omitempty"`
	TreatOpenedAsOutOfStock            int            `json:"treat_opened_as_out_of_stock,omitempty"`
	AutoReprintStockLabel              int            `json:"auto_reprint_stock_label,omitempty"`
	NoOwnStock                         int            `json:"no_own_stock,omitempty"`
	Userfields                         map[string]any `json:"userfields,omitempty"`
	ShouldNotBeFrozen                  int            `json:"should_not_be_frozen,omitempty"`
	DefaultConsumeLocationId           int            `json:"default_consume_location_id,omitempty"`
	MoveOnOpen                         int            `json:"move_on_open,omitempty"`
}

type ProductWithoutUserfields struct {
	Id                                 int       `json:"id,omitempty"`
	Name                               string    `json:"name,omitempty"`
	Description                        string    `json:"description,omitempty"`
	LocationId                         int       `json:"location_id,omitempty"`
	QuIdPurchase                       int       `json:"qu_id_purchase,omitempty"`
	QuIdStock                          int       `json:"qu_id_stock,omitempty"`
	EnableTareWeightHandling           int       `json:"enable_tare_weight_handling,omitempty"`
	NotCheckStockFulfillmentForRecipes int       `json:"not_check_stock_fulfillment_for_recipes,omitempty"`
	ProductGroupId                     int       `json:"product_group_id,omitempty"`
	TareWeight                         float64   `json:"tare_weight,omitempty"`
	MinStockAmount                     float64   `json:"min_stock_amount,omitempty"`
	DefaultBestBeforeDays              int       `json:"default_best_before_days,omitempty"`
	DefaultBestBeforeDaysAfterOpen     int       `json:"default_best_before_days_after_open,omitempty"`
	PictureFileName                    string    `json:"picture_file_name,omitempty"`
	RowCreatedTimestamp                time.Time `json:"row_created_timestamp"`
	ShoppingLocationId                 int       `json:"shopping_location_id,omitempty"`
	TreatOpenedAsOutOfStock            int       `json:"treat_opened_as_out_of_stock,omitempty"`
	AutoReprintStockLabel              int       `json:"auto_reprint_stock_label,omitempty"`
	NoOwnStock                         int       `json:"no_own_stock,omitempty"`
	ShouldNotBeFrozen                  int       `json:"should_not_be_frozen,omitempty"`
	DefaultConsumeLocationId           int       `json:"default_consume_location_id,omitempty"`
	MoveOnOpen                         int       `json:"move_on_open,omitempty"`
}

type QuantityUnit struct {
	Id                  int            `json:"id,omitempty"`
	Name                string         `json:"name,omitempty"`
	NamePlural          string         `json:"name_plural,omitempty"`
	Description         string         `json:"description,omitempty"`
	RowCreatedTimestamp time.Time      `json:"row_created_timestamp"`
	PluralForms         string         `json:"plural_forms,omitempty"`
	Userfields          map[string]any `json:"userfields,omitempty"`
}

type Location struct {
	Id                  int            `json:"id,omitempty"`
	Name                string         `json:"name,omitempty"`
	Description         string         `json:"description,omitempty"`
	RowCreatedTimestamp time.Time      `json:"row_created_timestamp"`
	Userfields          map[string]any `json:"userfields,omitempty"`
}

type ShoppingLocation struct {
	Id                  int            `json:"id,omitempty"`
	Name                string         `json:"name,omitempty"`
	Description         string         `json:"description,omitempty"`
	RowCreatedTimestamp time.Time      `json:"row_created_timestamp"`
	Userfields          map[string]any `json:"userfields,omitempty"`
}

type StockLocation struct {
	Id                int     `json:"id,omitempty"`
	ProductId         int     `json:"product_id,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
	LocationId        int     `json:"location_id,omitempty"`
	LocationName      string  `json:"location_name,omitempty"`
	LocationIsFreezer int     `json:"location_is_freezer,omitempty"`
}

type StockEntry struct {
	Id                  int       `json:"id,omitempty"`
	ProductId           int       `json:"product_id,omitempty"`
	LocationId          int       `json:"location_id,omitempty"`
	ShoppingLocationId  int       `json:"shopping_location_id,omitempty"`
	Amount              float64   `json:"amount,omitempty"`
	BestBeforeDate      time.Time `json:"best_before_date"`
	PurchasedDate       time.Time `json:"purchased_date"`
	StockId             string    `json:"stock_id,omitempty"`
	Price               float64   `json:"price,omitempty"`
	Open                int       `json:"open,omitempty"`
	OpenedDate          time.Time `json:"opened_date"`
	Note                string    `json:"note,omitempty"`
	RowCreatedTimestamp time.Time `json:"row_created_timestamp"`
}

type RecipeFulfillmentResponse struct {
	RecipeId                      int     `json:"recipe_id,omitempty"`
	NeedFulfilled                 bool    `json:"need_fulfilled,omitempty"`
	NeedFulfilledWithShoppingList bool    `json:"need_fulfilled_with_shopping_list,omitempty"`
	MissingProductsCount          int     `json:"missing_products_count,omitempty"`
	Costs                         float64 `json:"costs,omitempty"`
}

type ProductDetailsResponse struct {
	Product                           Product          `json:"product"`
	ProductBarcodes                   []ProductBarcode `json:"product_barcodes,omitempty"`
	QuantityUnitStock                 QuantityUnit     `json:"quantity_unit_stock"`
	DefaultQuantityUnitPurchase       QuantityUnit     `json:"default_quantity_unit_purchase"`
	DefaultQuantityUnitConsume        QuantityUnit     `json:"default_quantity_unit_consume"`
	QuantityUnitPrice                 QuantityUnit     `json:"quantity_unit_price"`
	LastPurchased                     time.Time        `json:"last_purchased"`
	LastUsed                          time.Time        `json:"last_used"`
	StockAmount                       float64          `json:"stock_amount,omitempty"`
	StockAmountOpened                 float64          `json:"stock_amount_opened,omitempty"`
	NextDueDate                       time.Time        `json:"next_due_date"`
	LastPrice                         float64          `json:"last_price,omitempty"`
	AvgPrice                          float64          `json:"avg_price,omitempty"`
	CurrentPrice                      float64          `json:"current_price,omitempty"`
	OldestPrice                       float64          `json:"oldest_price,omitempty"`
	LastShoppingLocationId            int              `json:"last_shopping_location_id,omitempty"`
	Location                          Location         `json:"location"`
	AverageShelfLifeDays              float64          `json:"average_shelf_life_days,omitempty"`
	SpoilRatePercent                  float64          `json:"spoil_rate_percent,omitempty"`
	HasChilds                         bool             `json:"has_childs,omitempty"`
	DefaultLocation                   Location         `json:"default_location"`
	QuConversionFactorPurchaseToStock float64          `json:"qu_conversion_factor_purchase_to_stock,omitempty"`
	QuConversionFactorPriceToStock    float64          `json:"qu_conversion_factor_price_to_stock,omitempty"`
}

type ProductPriceHistory struct {
	Date             time.Time        `json:"date"`
	Price            float64          `json:"price,omitempty"`
	ShoppingLocation ShoppingLocation `json:"shopping_location"`
}

type ProductBarcode struct {
	ProductId          int     `json:"product_id,omitempty"`
	Barcode            string  `json:"barcode,omitempty"`
	QuId               int     `json:"qu_id,omitempty"`
	ShoppingLocationId int     `json:"shopping_location_id,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	LastPrice          float64 `json:"last_price,omitempty"`
	Note               string  `json:"note,omitempty"`
}

type ExternalBarcodeLookupResponse struct {
	Name                    string  `json:"name,omitempty"`
	LocationId              int     `json:"location_id,omitempty"`
	QuIdPurchase            int     `json:"qu_id_purchase,omitempty"`
	QuIdStock               int     `json:"qu_id_stock,omitempty"`
	QuFactorPurchaseToStock float64 `json:"qu_factor_purchase_to_stock,omitempty"`
	Barcode                 string  `json:"barcode,omitempty"`
	Id                      int     `json:"id,omitempty"`
}

type ChoreDetailsResponse struct {
	Chore                          Chore     `json:"chore"`
	LastTracked                    time.Time `json:"last_tracked"`
	TrackCount                     int       `json:"track_count,omitempty"`
	LastDoneBy                     UserDto   `json:"last_done_by"`
	NextEstimatedExecutionTime     time.Time `json:"next_estimated_execution_time"`
	NextExecutionAssignedUser      UserDto   `json:"next_execution_assigned_user"`
	AverageExecutionFrequencyHours int       `json:"average_execution_frequency_hours,omitempty"`
}

type BatteryDetailsResponse struct {
	Chore                   Battery   `json:"chore"`
	LastCharged             time.Time `json:"last_charged"`
	ChargeCyclesCount       int       `json:"charge_cycles_count,omitempty"`
	NextEstimatedChargeTime time.Time `json:"next_estimated_charge_time"`
}

type Session struct {
	Id                  int       `json:"id,omitempty"`
	SessionKey          string    `json:"session_key,omitempty"`
	Expires             time.Time `json:"expires"`
	LastUsed            time.Time `json:"last_used"`
	RowCreatedTimestamp time.Time `json:"row_created_timestamp"`
}

type User struct {
	Id                  int       `json:"id,omitempty"`
	Username            string    `json:"username,omitempty"`
	FirstName           string    `json:"first_name,omitempty"`
	LastName            string    `json:"last_name,omitempty"`
	Password            string    `json:"password,omitempty"`
	PictureFileName     string    `json:"picture_file_name,omitempty"`
	RowCreatedTimestamp time.Time `json:"row_created_timestamp"`
}

type UserDto struct {
	Id                  int       `json:"id,omitempty"`
	Username            string    `json:"username,omitempty"`
	FirstName           string    `json:"first_name,omitempty"`
	LastName            string    `json:"last_name,omitempty"`
	DisplayName         string    `json:"display_name,omitempty"`
	PictureFileName     string    `json:"picture_file_name,omitempty"`
	RowCreatedTimestamp time.Time `json:"row_created_timestamp"`
}

type ApiKey struct {
	Id                  int       `json:"id,omitempty"`
	ApiKey              string    `json:"api_key,omitempty"`
	Expires             time.Time `json:"expires"`
	LastUsed            time.Time `json:"last_used"`
	RowCreatedTimestamp time.Time `json:"row_created_timestamp"`
}

type ShoppingListItem struct {
	Id                  int            `json:"id,omitempty"`
	ShoppingListId      int            `json:"shopping_list_id,omitempty"`
	ProductId           int            `json:"product_id,omitempty"`
	Note                string         `json:"note,omitempty"`
	Amount              float64        `json:"amount,omitempty"`
	RowCreatedTimestamp time.Time      `json:"row_created_timestamp"`
	Userfields          map[string]any `json:"userfields,omitempty"`
}

type Battery struct {
	Id                  int            `json:"id,omitempty"`
	Name                string         `json:"name,omitempty"`
	Description         string         `json:"description,omitempty"`
	UsedIn              string         `json:"used_in,omitempty"`
	ChargeIntervalDays  int            `json:"charge_interval_days,omitempty"`
	RowCreatedTimestamp time.Time      `json:"row_created_timestamp"`
	Userfields          map[string]any `json:"userfields,omitempty"`
}

type BatteryChargeCycleEntry struct {
	Id                  int       `json:"id,omitempty"`
	BatteryId           int       `json:"battery_id,omitempty"`
	TrackedTime         time.Time `json:"tracked_time"`
	RowCreatedTimestamp time.Time `json:"row_created_timestamp"`
}

type Chore struct {
	Id                                       int            `json:"id,omitempty"`
	Name                                     string         `json:"name,omitempty"`
	Description                              string         `json:"description,omitempty"`
	PeriodType                               string         `json:"period_type,omitempty"`
	PeriodConfig                             string         `json:"period_config,omitempty"`
	PeriodDays                               int            `json:"period_days,omitempty"`
	TrackDateOnly                            bool           `json:"track_date_only,omitempty"`
	Rollover                                 bool           `json:"rollover,omitempty"`
	AssignmentType                           string         `json:"assignment_type,omitempty"`
	AssignmentConfig                         string         `json:"assignment_config,omitempty"`
	NextExecutionAssignedToUserId            int            `json:"next_execution_assigned_to_user_id,omitempty"`
	StartDate                                time.Time      `json:"start_date"`
	RescheduledDate                          time.Time      `json:"rescheduled_date"`
	RescheduledNextExecutionAssignedToUserId int            `json:"rescheduled_next_execution_assigned_to_user_id,omitempty"`
	RowCreatedTimestamp                      time.Time      `json:"row_created_timestamp"`
	Userfields                               map[string]any `json:"userfields,omitempty"`
}

type ChoreLogEntry struct {
	Id                  int       `json:"id,omitempty"`
	ChoreId             int       `json:"chore_id,omitempty"`
	TrackedTime         time.Time `json:"tracked_time"`
	RowCreatedTimestamp time.Time `json:"row_created_timestamp"`
}

type StockLogEntry struct {
	Id                  int                  `json:"id,omitempty"`
	ProductId           int                  `json:"product_id,omitempty"`
	Amount              float64              `json:"amount,omitempty"`
	BestBeforeDate      time.Time            `json:"best_before_date"`
	PurchasedDate       time.Time            `json:"purchased_date"`
	UsedDate            time.Time            `json:"used_date"`
	Spoiled             bool                 `json:"spoiled,omitempty"`
	StockId             string               `json:"stock_id,omitempty"`
	TransactionId       string               `json:"transaction_id,omitempty"`
	TransactionType     StockTransactionType `json:"transaction_type,omitempty"`
	Note                string               `json:"note,omitempty"`
	RowCreatedTimestamp time.Time            `json:"row_created_timestamp"`
}

type StockJournal struct {
	CorrelationId       string               `json:"correlation_id,omitempty"`
	Undone              int                  `json:"undone,omitempty"`
	UndoneTimestamp     time.Time            `json:"undone_timestamp"`
	Amount              float64              `json:"amount,omitempty"`
	LocationId          int                  `json:"location_id,omitempty"`
	LocationName        string               `json:"location_name,omitempty"`
	ProductName         string               `json:"product_name,omitempty"`
	QuName              string               `json:"qu_name,omitempty"`
	QuNamePlural        string               `json:"qu_name_plural,omitempty"`
	UserDisplayName     string               `json:"user_display_name,omitempty"`
	Spoiled             bool                 `json:"spoiled,omitempty"`
	TransactionType     StockTransactionType `json:"transaction_type,omitempty"`
	RowCreatedTimestamp time.Time            `json:"row_created_timestamp"`
}

type StockJournalSummary struct {
	Amount          float64              `json:"amount,omitempty"`
	UserId          int                  `json:"user_id,omitempty"`
	ProductName     string               `json:"product_name,omitempty"`
	ProductId       int                  `json:"product_id,omitempty"`
	QuName          string               `json:"qu_name,omitempty"`
	QuNamePlural    string               `json:"qu_name_plural,omitempty"`
	UserDisplayName string               `json:"user_display_name,omitempty"`
	TransactionType StockTransactionType `json:"transaction_type,omitempty"`
}

type Error400 struct {
	ErrorMessage string `json:"error_message,omitempty"`
}

type Error500 struct {
	ErrorMessage string         `json:"error_message,omitempty"`
	ErrorDetails map[string]any `json:"error_details,omitempty"`
}

type CurrentStockResponse struct {
	ProductId              int                      `json:"product_id,omitempty"`
	Amount                 float64                  `json:"amount,omitempty"`
	AmountAggregated       float64                  `json:"amount_aggregated,omitempty"`
	AmountOpened           float64                  `json:"amount_opened,omitempty"`
	AmountOpenedAggregated float64                  `json:"amount_opened_aggregated,omitempty"`
	BestBeforeDate         time.Time                `json:"best_before_date"`
	IsAggregatedAmount     bool                     `json:"is_aggregated_amount,omitempty"`
	Product                ProductWithoutUserfields `json:"product"`
}

type CurrentChoreResponse struct {
	ChoreId                       int       `json:"chore_id,omitempty"`
	ChoreName                     string    `json:"chore_name,omitempty"`
	LastTrackedTime               time.Time `json:"last_tracked_time"`
	TrackDateOnly                 bool      `json:"track_date_only,omitempty"`
	NextEstimatedExecutionTime    time.Time `json:"next_estimated_execution_time"`
	NextExecutionAssignedToUserId int       `json:"next_execution_assigned_to_user_id,omitempty"`
	IsRescheduled                 bool      `json:"is_rescheduled,omitempty"`
	IsReassigned                  bool      `json:"is_reassigned,omitempty"`
	NextExecutionAssignedUser     UserDto   `json:"next_execution_assigned_user"`
}

type CurrentBatteryResponse struct {
	BatteryId               int       `json:"battery_id,omitempty"`
	LastTrackedTime         time.Time `json:"last_tracked_time"`
	NextEstimatedChargeTime time.Time `json:"next_estimated_charge_time"`
}

type CurrentVolatilStockResponse struct {
	DueProducts     []CurrentStockResponse `json:"due_products,omitempty"`
	OverdueProducts []CurrentStockResponse `json:"overdue_products,omitempty"`
	ExpiredProducts []CurrentStockResponse `json:"expired_products,omitempty"`
	MissingProducts []any                  `json:"missing_products,omitempty"`
}

type Task struct {
	Id                  int            `json:"id,omitempty"`
	Name                string         `json:"name,omitempty"`
	Description         string         `json:"description,omitempty"`
	DueDate             time.Time      `json:"due_date"`
	Done                int            `json:"done,omitempty"`
	DoneTimestamp       time.Time      `json:"done_timestamp"`
	CategoryId          int            `json:"category_id,omitempty"`
	AssignedToUserId    int            `json:"assigned_to_user_id,omitempty"`
	RowCreatedTimestamp time.Time      `json:"row_created_timestamp"`
	Userfields          map[string]any `json:"userfields,omitempty"`
}

type TaskCategory struct {
	Id                  int       `json:"id,omitempty"`
	Name                string    `json:"name,omitempty"`
	Description         string    `json:"description,omitempty"`
	RowCreatedTimestamp time.Time `json:"row_created_timestamp"`
}

type CurrentTaskResponse struct {
	Id                  int          `json:"id,omitempty"`
	Name                string       `json:"name,omitempty"`
	Description         string       `json:"description,omitempty"`
	DueDate             time.Time    `json:"due_date"`
	Done                int          `json:"done,omitempty"`
	DoneTimestamp       time.Time    `json:"done_timestamp"`
	CategoryId          int          `json:"category_id,omitempty"`
	AssignedToUserId    int          `json:"assigned_to_user_id,omitempty"`
	RowCreatedTimestamp time.Time    `json:"row_created_timestamp"`
	AssignedToUser      UserDto      `json:"assigned_to_user"`
	Category            TaskCategory `json:"category"`
}

type DbChangedTimeResponse struct {
	ChangedTime time.Time `json:"changed_time"`
}

type TimeResponse struct {
	Timezone         string    `json:"timezone,omitempty"`
	TimeLocal        time.Time `json:"time_local"`
	TimeLocalSqlite3 time.Time `json:"time_local_sqlite3"`
	TimeUtc          time.Time `json:"time_utc"`
	Timestamp        int       `json:"timestamp,omitempty"`
	Offset           int       `json:"offset,omitempty"`
}

type UserSetting struct {
	Value string `json:"value,omitempty"`
}

type MissingLocalizationRequest struct {
	Text string `json:"text,omitempty"`
}
