package bigcommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// Order describes the product resource
type Order struct {
	ID                   int           `json:"id"`
	CustomerID           int           `json:"customer_id"`
	DateCreated          BCTime        `json:"date_created"`
	DateModified         BCTime        `json:"date_modified"`
	DateShipped          BCTime        `json:"date_shipped"`
	StatusID             int           `json:"status_id"`
	Status               string        `json:"status"`
	HandlingCostExTax    float64       `json:"handling_cost_ex_tax,string"`
	HandlingCostIncTax   float64       `json:"handling_cost_inc_tax,string"`
	HandlingCostTax      float64       `json:"handling_cost_tax,string"`
	ShippingCostExTax    float64       `json:"shipping_cost_ex_tax,string"`
	ShippingCostIncTax   float64       `json:"shipping_cost_inc_tax,string"`
	ShippingCostTax      float64       `json:"shipping_cost_tax,string"`
	SubTotalExTax        float64       `json:"subtotal_ex_tax,string"`
	SubTotalIncTax       float64       `json:"subtotal_inc_tax,string"`
	SubTotalTax          float64       `json:"subtotal_tax,string"`
	TotalExTax           float64       `json:"total_ex_tax,string"`
	TotalIncTax          float64       `json:"total_inc_tax,string"`
	TotalTax             float64       `json:"total_tax,string"`
	BaseShippingCost     float64       `json:"base_shipping_cost,string"`
	ItemsTotal           int           `json:"items_total"`
	PaymentMethod        string        `json:"payment_method"`
	PaymentStatus        string        `json:"payment_status"`
	IPAddress            string        `json:"ip_address"`
	CurrencyID           int           `json:"currency_id"`
	CurrencyCode         string        `json:"currency_code"`
	StaffNotes           string        `json:"staff_notes"`
	CustomerMessage      string        `json:"customer_message"`
	DiscountAmount       string        `json:"discount_amount"`
	CouponDiscount       string        `json:"counpon_discount"`
	ShippingAddressCount int           `json:"shipping_address_count"`
	BillingAddress       AddressEntity `json:"billing_address"`
}

// OrderService adds the APIs for the Order resource.
type OrderService struct {
	sling      *sling.Sling
	httpClient *http.Client
}

func newOrderService(sling *sling.Sling, httpClient *http.Client) *OrderService {
	return &OrderService{
		sling:      sling.Path("orders/"),
		httpClient: httpClient,
	}
}

// OrderListParams are the parameters for OrderService.List
type OrderListParams struct {
	Page          int     `url:"page,omitempty"`
	Limit         int     `url:"limit,omitempty"`
	Sort          string  `url:"sort,omitempty"`
	MinID         int     `url:"min_id,omitempty"`
	MaxID         int     `url:"max_id,omitempty"`
	MinTotal      float64 `url:"min_total,omitempty"`
	MaxTotal      float64 `url:"max_total,omitempty"`
	CustomerID    *int    `url:"customer_id,omitempty"`
	Email         string  `url:"email,omitempty"`
	StatusID      *int    `url:"status_id,omitempty"`
	PaymentMethod string  `url:"payment_method,omitempty"`
	//TODO: add date and boolean based params.
}

// List returns a list of Orders matching the given OrderListParams.
func (s *OrderService) List(ctx context.Context, params *OrderListParams) ([]Order, *http.Response, error) {
	var orders []Order
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().QueryStruct(params), s.httpClient, &orders, apiError)
	return orders, resp, relevantError(err, *apiError)
}

// Count returns an OrderCount for Orders that matches the given OrderListParams.
func (s *OrderService) Count(ctx context.Context, params *OrderListParams) (int, *http.Response, error) {
	var count count
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.Get("count").QueryStruct(params), s.httpClient, &count, apiError)
	return count.Count, resp, relevantError(err, *apiError)
}

// Show returns the requested Order.
func (s *OrderService) Show(ctx context.Context, id int32) (*Order, *http.Response, error) {
	order := new(Order)
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().Get(fmt.Sprintf("%d", id)), s.httpClient, order, apiError)
	return order, resp, relevantError(err, *apiError)
}

// OrderProduct defines a product to be included in the OrderBody.
// Regular Products require: ProductID and Quantity
// Custom Products require: Name, Quantity and PriceIncTax / PriceExTax
type OrderProduct struct {
	ProductID   int     `json:"product_id,omitempty"`
	ProductName string  `json:"name,omitempty"`
	Quantity    int     `json:"quantity"`
	PriceIncTax float64 `json:"price_inc_tax,omitempty"`
	PriceExTax  float64 `json:"price_ex_tax,omitempty"`
}

// OrderBody describes the order information given when creating a new Order.
type OrderBody struct {
	ExternalSource     string          `json:"external_source"`
	CustomerID         *int            `json:"customer_id"`
	StatusID           *int            `json:"status_id"`
	BillingAddress     AddressEntity   `json:"billing_address"`
	Products           []OrderProduct  `json:"products"`
	ShippingCostIncTax float64         `json:"shipping_cost_inc_tax,omitempty"`
	ShippingCostExTax  float64         `json:"shipping_cost_ex_tax,omitempty"`
	HandlingCostIncTax float64         `json:"handling_cost_inc_tax,omitempty"`
	HandlingCostExTax  float64         `json:"handling_cost_ex_tax,omitempty"`
	ShippingAddresses  AddressEntities `json:"shipping_addresses,omitempty"`
	CustomerMessage    string          `json:"customer_message"`
	StaffNotes         string          `json:"staff_notes"`
	PaymentMethod      string          `json:"payment_method"`
}

// New creates a new Order with the specified information and returns the new order.
func (s *OrderService) New(ctx context.Context, body *OrderBody) (*Order, *http.Response, error) {
	order := new(Order)
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().Post("").BodyJSON(body), s.httpClient, order, apiError)
	return order, resp, relevantError(err, *apiError)
}

// OrderEditParams describes the fields that are editable on an Order.
type OrderEditParams struct {
	CustomerID      *int           `json:"customer_id,omitempty"`
	StatusID        *int           `json:"status_id,omitempty"`
	IPAddress       string         `json:"ip_address,omitempty"`
	StaffNotes      string         `json:"staff_notes,omitempty"`
	CustomerMessage string         `json:"customer_message,omitempty"`
	BillingAddress  *AddressEntity `json:"billing_address,omitempty"`
}

// Edit updates the given OrderEditParams of the given Order.
func (s *OrderService) Edit(ctx context.Context, id int, params *OrderEditParams) (*Order, *http.Response, error) {
	order := new(Order)
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().Put(fmt.Sprintf("%d", id)).BodyJSON(params), s.httpClient, order, apiError)
	return order, resp, relevantError(err, *apiError)
}
