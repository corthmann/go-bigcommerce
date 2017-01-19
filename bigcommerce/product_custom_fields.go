package bigcommerce

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
	"golang.org/x/net/context"
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
	sling      *sling.Sling
	httpClient *http.Client
}

func newProductCustomFieldService(sling *sling.Sling, httpClient *http.Client) *ProductCustomFieldService {
	return &ProductCustomFieldService{
		sling:      sling.Path("products/"),
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

	resp, err := performRequest(ctx, s.sling.New().Get(fmt.Sprintf("%d/custom_fields", productID)).QueryStruct(params), s.httpClient, &customFields, apiError)
	return customFields, resp, relevantError(err, *apiError)
}

// Show returns the requested ProductCustomField.
func (s *ProductCustomFieldService) Show(ctx context.Context, productID int, id int) (*ProductCustomField, *http.Response, error) {
	customField := new(ProductCustomField)
	apiError := new(APIError)

	resp, err := performRequest(ctx, s.sling.New().Get(fmt.Sprintf("%d/custom_fields/%d", productID, id)), s.httpClient, customField, apiError)
	return customField, resp, relevantError(err, *apiError)
}
