package bigcommerce

import (
	"fmt"
	"net/http"

	"context"
)

const productServicePath = "products/"

// Product describes the product resource
type Product struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	Sku            string             `json:"sku"`
	Description    string             `json:"description"`
	Price          string             `json:"price"`
	CostPrice      string             `json:"cost_price"`
	RetailPrice    string             `json:"retail_price"`
	InventoryLevel int                `json:"inventory_level"`
	TotalSold      int                `json:"total_sold"`
	Availability   string             `json:"availability"`
	PrimaryImage   PrimaryImageEntity `json:"primary_image"`
}

// PrimaryImageEntity describes the image entity.
type PrimaryImageEntity struct {
	StandardURL  string `json:"standard_url"`
	TinyURL      string `json:"tiny_url"`
	ThumbnailURL string `json:"thumbnail_url"`
	ZoomURL      string `json:"zoom_url"`
}

// ProductService adds the APIs for the Product resource.
type ProductService struct {
	config     *ClientConfig
	httpClient *http.Client
}

func newProductService(config *ClientConfig, httpClient *http.Client) *ProductService {
	return &ProductService{
		config:     config,
		httpClient: httpClient,
	}
}

// ProductListParams are the parameters for ProductService.List
type ProductListParams struct {
	MinID             int    `url:"min_id,omitempty"`
	MaxID             int    `url:"max_id,omitempty"`
	Name              string `url:"name,omitempty"`
	Sku               string `url:"sku,omitempty"`
	Availability      string `url:"availability,omitempty"`
	IsVisible         string `url:"is_visible,omitempty"`
	IsFeatured        string `url:"is_featured,omitempty"`
	MinInventoryLevel int    `url:"min_inventory_level,omitempty"`
	MaxInventoryLevel int    `url:"max_inventory_level,omitempty"`
}

// List returns a list of Products matching the given ProductListParams.
func (s *ProductService) List(ctx context.Context, params *ProductListParams) ([]Product, *http.Response, error) {
	var products []Product
	var apiError APIError

	response, err := performGET(ctx, s.httpClient, s.config, productServicePath, params, &products, apiError)

	return products, response, relevantError(err, apiError)
}

// Show returns the requested Product.
func (s *ProductService) Show(ctx context.Context, id int32) (*Product, *http.Response, error) {
	product := new(Product)
	apiError := new(APIError)

	path := fmt.Sprintf("%v%v", productServicePath, id)
	response, err := performGET(ctx, s.httpClient, s.config, path, nil, &product, apiError)

	return product, response, relevantError(err, *apiError)
}
