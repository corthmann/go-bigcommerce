package bigcommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// OrderStatus describes the product resource
type OrderStatus struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Order int32  `json:"order"`
}

// OrderStatusService adds the APIs for the Product resource.
type OrderStatusService struct {
	sling      *sling.Sling
	httpClient *http.Client
}

func newOrderStatusService(sling *sling.Sling, httpClient *http.Client) *OrderStatusService {
	return &OrderStatusService{
		sling:      sling.Path("order_statuses/"),
		httpClient: httpClient,
	}
}

// OrderStatusListParams are the parameters for OrderStatusService.List
type OrderStatusListParams struct {
	Page  int32 `url:"page,omitempty"`
	Limit int32 `url:"limit,omitempty"`
}

// List returns a list of Products matching the given ProductListParams.
func (s *OrderStatusService) List(ctx context.Context, params *OrderStatusListParams) ([]OrderStatus, *http.Response, error) {
	var os []OrderStatus
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().QueryStruct(params), s.httpClient, &os, apiError)
	return os, resp, relevantError(err, *apiError)
}

// Show returns the requested OrderStatus.
func (s *OrderStatusService) Show(ctx context.Context, id int32) (*OrderStatus, *http.Response, error) {
	orderStatus := new(OrderStatus)
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().Get(fmt.Sprintf("%d", id)), s.httpClient, orderStatus, apiError)
	return orderStatus, resp, relevantError(err, *apiError)
}
