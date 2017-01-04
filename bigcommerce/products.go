package bigcommerce

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// Products defines a list of the Product object.
type Products []Product

// Product describes the product resource
type Product struct {
	ID             int32              `json:"id"`
	Name           string             `json:"name"`
	Sku            string             `json:"sku"`
	Description    string             `json:"description"`
	Price          string             `json:"price"`
	CostPrice      string             `json:"cost_price"`
	RetailPrice    string             `json:"retail_price"`
	InventoryLevel int32              `json:"inventory_level"`
	TotalSold      int32              `json:"total_sold"`
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
	sling *sling.Sling
}

func newProductService(sling *sling.Sling) *ProductService {
	return &ProductService{
		sling: sling.Path("products/"),
	}
}

// ProductListParams are the parameters for ProductService.List
type ProductListParams struct {
	MinID             int32  `url:"min_id,omitempty"`
	MaxID             int32  `url:"max_id,omitempty"`
	Name              string `url:"name,omitempty"`
	Sku               string `url:"sku,omitempty"`
	Availability      string `url:"availability,omitempty"`
	IsVisible         string `url:"is_visible,omitempty"`
	IsFeatured        string `url:"is_featured,omitempty"`
	MinInventoryLevel int32  `url:"min_inventory_level,omitempty"`
	MaxInventoryLevel int32  `url:"max_inventory_level,omitempty"`
}

// List returns a list of Products matching the given ProductListParams.
func (s *ProductService) List(params *ProductListParams) (*Products, *http.Response, error) {
	products := new(Products)
	apiError := new(APIError)

	resp, err := s.sling.New().QueryStruct(params).Receive(products, apiError)
	return products, resp, relevantError(err, *apiError)
}

// Show returns the requested Product.
func (s *ProductService) Show(id int32) (*Product, *http.Response, error) {
	product := new(Product)
	apiError := new(APIError)

	resp, err := s.sling.New().Get(fmt.Sprintf("%d", id)).Receive(product, apiError)
	return product, resp, relevantError(err, *apiError)
}
