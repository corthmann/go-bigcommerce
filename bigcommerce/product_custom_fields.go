package bigcommerce

import (
	"fmt"
	"net/http"

	"context"
)

// ProductCustomField describes the product custom field resource
type ProductCustomField struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Text      string `json:"text"`
}

// ProductCustomFieldService adds the APIs for the ProductCustomField resource.
type ProductCustomFieldService struct {
	config     *ClientConfig
	httpClient *http.Client
}

func newProductCustomFieldService(config *ClientConfig, httpClient *http.Client) *ProductCustomFieldService {
	return &ProductCustomFieldService{
		config:     config,
		httpClient: httpClient,
	}
}

// ProductCustomFieldListParams are the parameters for ProductCustomFieldService.List
type ProductCustomFieldListParams struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// List returns a list of ProductCustomFields matching the given ProductCustomFieldListParams.
func (s *ProductCustomFieldService) List(ctx context.Context, productID int, params *ProductCustomFieldListParams) ([]ProductCustomField, *http.Response, error) {
	var customFields []ProductCustomField
	apiError := new(APIError)

	response, err := performGET(ctx, s.httpClient, s.config, s.servicePath(productID), params, &customFields, apiError)

	return customFields, response, relevantError(err, *apiError)
}

// Show returns the requested ProductCustomField.
func (s *ProductCustomFieldService) Show(ctx context.Context, productID int, id int) (*ProductCustomField, *http.Response, error) {
	customField := new(ProductCustomField)
	apiError := new(APIError)

	path := fmt.Sprintf("%v/%d", s.servicePath(productID), id)
	response, err := performGET(ctx, s.httpClient, s.config, path, nil, &customField, apiError)

	return customField, response, relevantError(err, *apiError)
}

func (s *ProductCustomFieldService) servicePath(productID int) string {
	return fmt.Sprintf("products/%d/custom_fields", productID)
}
