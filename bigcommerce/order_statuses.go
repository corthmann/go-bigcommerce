package bigcommerce

import (
	"context"
	"fmt"
	"net/http"
)

const orderStatusServicePath = "order_statuses/"

// OrderStatus describes the product resource
type OrderStatus struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

// OrderStatusService adds the APIs for the Product resource.
type OrderStatusService struct {
	config     *ClientConfig
	httpClient *http.Client
}

func newOrderStatusService(config *ClientConfig, httpClient *http.Client) *OrderStatusService {
	return &OrderStatusService{
		config:     config,
		httpClient: httpClient,
	}
}

// OrderStatusListParams are the parameters for OrderStatusService.List
type OrderStatusListParams struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// List returns a list of Products matching the given ProductListParams.
func (s *OrderStatusService) List(ctx context.Context, params *OrderStatusListParams) ([]OrderStatus, *http.Response, error) {
	var os []OrderStatus
	apiError := new(APIError)

	response, err := performGET(ctx, s.httpClient, s.config, orderStatusServicePath, params, &os, apiError)

	return os, response, relevantError(err, *apiError)
}

// Show returns the requested OrderStatus.
func (s *OrderStatusService) Show(ctx context.Context, id int) (*OrderStatus, *http.Response, error) {
	orderStatus := new(OrderStatus)
	apiError := new(APIError)

	path := fmt.Sprintf("%v%v", orderStatusServicePath, id)
	response, err := performGET(ctx, s.httpClient, s.config, path, nil, &orderStatus, apiError)

	return orderStatus, response, relevantError(err, *apiError)
}
