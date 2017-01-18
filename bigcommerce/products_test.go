package bigcommerce

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/products/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"sku": "VIV-300020"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{ "id": 123 }]`)
	})

	expected := []Product{
		{ID: 123},
	}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	params := &ProductListParams{
		Sku: "VIV-300020",
	}
	products, _, err := client.Products.List(context.Background(), params)
	assert.Nil(t, err)
	assert.Equal(t, expected, products)
}

func TestProductService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/products/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{ "id": 123 }`)
	})

	expected := &Product{ID: 123}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	product, _, err := client.Products.Show(context.Background(), 123)
	assert.Nil(t, err)
	assert.Equal(t, expected, product)
}
