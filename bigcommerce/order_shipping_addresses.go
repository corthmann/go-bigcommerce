package bigcommerce

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// OrderShippingAddress describes how shipping addresses are returned for an order.
// It contains an ID, OrderID and an AddressEntity.
type OrderShippingAddress struct {
	AddressEntity
	ID      int `json:"id"`
	OrderID int `json:"order_id"`
}

// OrderShippingAddressService adds the APIs for the OrderShippingAddress resource.
type OrderShippingAddressService struct {
	config     *ClientConfig
	httpClient *http.Client
}

func newOrderShippingAddressService(config *ClientConfig, httpClient *http.Client) *OrderShippingAddressService {
	return &OrderShippingAddressService{
		config:     config,
		httpClient: httpClient,
	}
}

// OrderShippingAddressListParams are the parameters for OrderShippingAddressService.List
type OrderShippingAddressListParams struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// List returns a list of OrderShippingAddresses matching the given OrderShippingAddressListParams.
func (s *OrderShippingAddressService) List(ctx context.Context, orderID int, params *OrderShippingAddressListParams) ([]OrderShippingAddress, *http.Response, error) {
	var osa []OrderShippingAddress
	var apiError APIError

	response, err := performGET(ctx, s.httpClient, s.config, s.servicePath(orderID), params, &osa, &apiError)

	return osa, response, relevantError(err, apiError)
}

// Count returns an OrderShippingAddressCount for OrderShippingAddresses that matches the given OrderShippingAddressListParams.
func (s *OrderShippingAddressService) Count(ctx context.Context, orderID int, params *OrderShippingAddressListParams) (int, *http.Response, error) {
	var cnt count
	var apiError APIError

	path := strings.Join([]string{s.servicePath(orderID), "count"}, "")
	response, err := performGET(ctx, s.httpClient, s.config, path, params, &cnt, &apiError)

	return cnt.Count, response, relevantError(err, apiError)
}

// Show returns the requested OrderShippingAddress.
func (s *OrderShippingAddressService) Show(ctx context.Context, orderID int, id int) (*OrderShippingAddress, *http.Response, error) {
	osa := new(OrderShippingAddress)
	var apiError APIError

	path := fmt.Sprintf("%v%v", s.servicePath(orderID), id)
	response, err := performGET(ctx, s.httpClient, s.config, path, nil, &osa, &apiError)

	return osa, response, relevantError(err, apiError)
}

func (s *OrderShippingAddressService) servicePath(orderID int) string {
	return fmt.Sprintf("orders/%d/shipping_addresses/", orderID)
}
