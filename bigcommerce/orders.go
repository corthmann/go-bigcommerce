package bigcommerce

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// Orders defines a list of the Order object.
type Orders []Order

// Order describes the product resource
type Order struct {
	ID                   int32         `json:"id"`
	CustomerID           int32         `json:"customer_id"`
	DateCreated          string        `json:"date_created"`
	DateModified         string        `json:"date_modified"`
	DateShipped          string        `json:"date_shipped"`
	StatusID             int32         `json:"status_id"`
	Status               string        `json:"status"`
	HandlingCostExTax    string        `json:"handling_cost_ex_tax"`
	HandlingCostIncTax   string        `json:"handling_cost_inc_tax"`
	HandlingCostTax      string        `json:"handling_cost_tax"`
	ShippingCostExTax    string        `json:"shipping_cost_ex_tax"`
	ShippingCostIncTax   string        `json:"shipping_cost_inc_tax"`
	ShippingCostTax      string        `json:"shipping_cost_tax"`
	SubTotalExTax        string        `json:"subtotal_ex_tax"`
	SubTotalIncTax       string        `json:"subtotal_inc_tax"`
	SubTotalTax          string        `json:"subtotal_tax"`
	TotalExTax           string        `json:"total_ex_tax"`
	TotalIncTax          string        `json:"total_inc_tax"`
	TotalTax             string        `json:"total_tax"`
	BaseShippingCost     string        `json:"base_shipping_cost"`
	ItemsTotal           int32         `json:"items_total"`
	PaymentMethod        string        `json:"payment_method"`
	PaymentStatus        string        `json:"payment_status"`
	IPAddress            string        `json:"ip_address"`
	CurrencyID           int32         `json:"currency_id"`
	CurrencyCode         string        `json:"currency_code"`
	StaffNotes           string        `json:"staff_notes"`
	CustomerMessage      string        `json:"customer_message"`
	DiscountAmount       string        `json:"discount_amount"`
	CouponDiscount       string        `json:"counpon_discount"`
	ShippingAddressCount int32         `json:"shipping_address_count"`
	BillingAddress       AddressEntity `json:"billing_address"`
}

// AddressEntity describes the address entity.
type AddressEntity struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Company     string `json:"company"`
	Street1     string `json:"street_1"`
	Street2     string `json:"street_2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	CountryIso2 string `json:"country_iso2"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

// OrderService adds the APIs for the Order resource.
type OrderService struct {
	sling *sling.Sling
}

func newOrderService(sling *sling.Sling) *OrderService {
	return &OrderService{
		sling: sling.Path("orders/"),
	}
}

// OrderListParams are the parameters for OrderService.List
type OrderListParams struct {
	MinID         int32   `url:"min_id,omitempty"`
	MaxID         int32   `url:"max_id,omitempty"`
	MinTotal      float32 `url:"min_total,omitempty"`
	MaxTotal      float32 `url:"max_total,omitempty"`
	CustomerID    string  `url:"customer_id,omitempty"` // this is actually an int, but doesn't come into the url when it is = zero.... which is bad..
	Email         string  `url:"email,omitempty"`
	StatusID      string  `url:"status_id,omitempty"`
	PaymentMethod string  `url:"payment_method,omitempty"`
	//TODO: add date and boolean based params.
}

// List returns a list of Orders matching the given OrderListParams.
func (s *OrderService) List(params *OrderListParams) (*Orders, *http.Response, error) {
	orders := new(Orders)
	apiError := new(APIError)

	resp, err := s.sling.New().QueryStruct(params).Receive(orders, apiError)
	return orders, resp, relevantError(err, *apiError)
}

// Show returns the requested Order.
func (s *OrderService) Show(id int32) (*Order, *http.Response, error) {
	order := new(Order)
	apiError := new(APIError)

	resp, err := s.sling.New().Get(fmt.Sprintf("%d", id)).Receive(order, apiError)
	return order, resp, relevantError(err, *apiError)
}
