package bigcommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// Order describes the product resource
type OrderShippingAddress struct {
	AddressEntity
	ID      int32 `json:"id"`
	OrderID int32 `json:"order_id"`
}

// OrderShippingAddressService adds the APIs for the OrderShippingAddress resource.
type OrderShippingAddressService struct {
	sling      *sling.Sling
	httpClient *http.Client
}

func newOrderShippingAddressService(sling *sling.Sling, httpClient *http.Client) *OrderShippingAddressService {
	return &OrderShippingAddressService{
		sling:      sling.Path("orders/"),
		httpClient: httpClient,
	}
}

// OrderShippingAddressListParams are the parameters for OrderShippingAddressService.List
type OrderShippingAddressListParams struct {
	Page  int32 `url:"page,omitempty"`
	Limit int32 `url:"limit,omitempty"`
}

// List returns a list of OrderShippingAddresses matching the given OrderShippingAddressListParams.
func (s *OrderShippingAddressService) List(ctx context.Context, orderID int32, params *OrderShippingAddressListParams) ([]OrderShippingAddress, *http.Response, error) {
	var osa []OrderShippingAddress
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().Get(fmt.Sprintf("%d/shipping_addresses", orderID)).QueryStruct(params), s.httpClient, &osa, apiError)
	return osa, resp, relevantError(err, *apiError)
}

// Count returns an OrderShippingAddressCount for OrderShippingAddresses that matches the given OrderShippingAddressListParams.
func (s *OrderShippingAddressService) Count(ctx context.Context, orderID int32, params *OrderShippingAddressListParams) (int, *http.Response, error) {
	var cnt count
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().Get(fmt.Sprintf("%d/shipping_addresses/count", orderID)).QueryStruct(params), s.httpClient, &cnt, apiError)
	return cnt.Count, resp, relevantError(err, *apiError)
}

// Show returns the requested OrderShippingAddress.
func (s *OrderShippingAddressService) Show(ctx context.Context, orderID int32, id int32) (*OrderShippingAddress, *http.Response, error) {
	osa := new(OrderShippingAddress)
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().Get(fmt.Sprintf("%d/shipping_addresses/%d", orderID, id)), s.httpClient, osa, apiError)
	return osa, resp, relevantError(err, *apiError)
}
