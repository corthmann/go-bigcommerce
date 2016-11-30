package bigcommerce

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Products struct {
	Products []Product
}

// type Products []map[Product]Product
// type Products struct {
// 	Products map[Product]Product
// }

// type Products struct {
// 	Products []Product {
// 		data map[Product]Product
// 	}
// }

type Product struct {
	Id             int32              `json:"id"`
	Name           string             `json:"name"`
	Sku            string             `json:"name"`
	Description    string             `json:"description"`
	Price          string             `json:"price"`
	CostPrice      string             `json:"cost_price"`
	RetailPrice    string             `json:"retail_price"`
	InventoryLevel int32              `json:"inventory_level"`
	TotalSold      int32              `json:"total_sold"`
	Availability   string             `json:"availability"`
	PrimaryImage   PrimaryImageEntity `json:"primary_image"`
}

type PrimaryImageEntity struct {
	StandardUrl  string `json:"standard_url"`
	TinyUrl      string `json:"tiny_url"`
	ThumbnailUrl string `json:"thumbnail_url"`
	ZoomUrl      string `json:"zoom_url"`
}

type ProductService struct {
	sling *sling.Sling
}

func newProductService(sling *sling.Sling) *ProductService {
	return &ProductService{
		sling: sling.Path("products"),
	}
}

// ProductListParams are the parameters for ProductService.List
type ProductListParams struct {
	MinId             int32  `url:"min_id,omitempty"`
	MaxId             int32  `url:"max_id,omitempty"`
	Name              string `url:"name,omitempty"`
	Sku               string `url:"sku,omitempty"`
	Availability      string `url:"availability,omitempty"`
	IsVisible         string `url:"is_visible,omitempty"`
	IsFeatured        string `url:"is_featured,omitempty"`
	MinInventoryLevel int32  `url:"min_inventory_level,omitempty"`
	MaxInventoryLevel int32  `url:"max_inventory_level,omitempty"`
}

func (s *ProductService) List(params *ProductListParams) (*Products, *http.Response, error) {

	// // text := "[{\"Id\": 100, \"Name\": \"Go\"}, {\"Id\": 200, \"Name\": \"Java\"}]"
	// // Get byte slice from string.
	// // bytes := []byte(Products)
	//
	// // Unmarshal string into structs.
	// var products Products
	// json.Unmarshal(bytes, &products)
	products := new(Products)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(params).Receive(products, apiError)
	// fmt.Println(resp)
	// fmt.Println(apiError)
	// fmt.Println(products)
	// fmt.Println(s.sling.New().QueryStruct(params))
	return products, resp, relevantError(err, *apiError)
}
