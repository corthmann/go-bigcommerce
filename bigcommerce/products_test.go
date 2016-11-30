package bigcommerce

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductService_List(t *testing.T) {
	// httpClient, mux, server := testServer()
	httpClient, _, server := testServer()
	defer server.Close()

	// mux.HandleFunc("/1.1/products", func(w http.ResponseWriter, r *http.Request) {
	// 	assertMethod(t, "GET", r)
	// 	assertQuery(t, map[string]string{"screen_name": "dghubble", "count": "5", "cursor": "1516933260114270762", "skip_status": "true", "include_user_entities": "false"}, r)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	fmt.Fprintf(w, `{"users": [{"id": 123}], "next_cursor":1516837838944119498,"next_cursor_str":"1516837838944119498","previous_cursor":-1516924983503961435,"previous_cursor_str":"-1516924983503961435"}`)
	// })
	// expected := &Products{
	// 	[]Product{Product{Id: 123}},
	// }
	//
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "",
		UserName: "",
		Password: ""})
	fmt.Println(client.sling)
	params := &ProductListParams{
		Sku: "VIV-300020",
	}
	products, _, err := client.Products.List(params)
	fmt.Println(products)
	assert.Nil(t, err)
	// assert.Equal(t, expected, products)
}
