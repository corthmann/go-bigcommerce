package bigcommerce

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// OrderStatuses defines a list of the OrderStatus object.
type OrderStatuses []OrderStatus

// OrderStatus describes the product resource
type OrderStatus struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Order int32  `json:"order"`
}

// OrderStatusService adds the APIs for the Product resource.
type OrderStatusService struct {
	sling *sling.Sling
}

func newOrderStatusService(sling *sling.Sling) *OrderStatusService {
	return &OrderStatusService{
		sling: sling.Path("order_statuses/"),
	}
}

// OrderStatusListParams are the parameters for OrderStatusService.List
type OrderStatusListParams struct {
	Page  int32 `url:"page,omitempty"`
	Limit int32 `url:"limit,omitempty"`
}

// List returns a list of Products matching the given ProductListParams.
func (s *OrderStatusService) List(params *OrderStatusListParams) (*OrderStatuses, *http.Response, error) {
	orderStatuses := new(OrderStatuses)
	apiError := new(APIError)

	resp, err := s.sling.New().QueryStruct(params).Receive(orderStatuses, apiError)
	return orderStatuses, resp, relevantError(err, *apiError)
}

// Show returns the requested OrderStatus.
func (s *OrderStatusService) Show(id int32) (*OrderStatus, *http.Response, error) {
	orderStatus := new(OrderStatus)
	apiError := new(APIError)

	resp, err := s.sling.New().Get(fmt.Sprintf("%d", id)).Receive(orderStatus, apiError)
	return orderStatus, resp, relevantError(err, *apiError)
}
