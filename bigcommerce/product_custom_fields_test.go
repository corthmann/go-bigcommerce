package bigcommerce

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductCustomFieldService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/products/12/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"page": "1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{ "id": 123 }]`)
	})

	expected := []ProductCustomField{
		{ID: 123},
	}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	params := &ProductCustomFieldListParams{
		Page: 1,
	}
	customFields, _, err := client.ProductCustomFields.List(context.Background(), 12, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, customFields)
}

func TestProductCustomFieldService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/products/12/custom_fields/3", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{ "id": 123 }`)
	})

	expected := &ProductCustomField{ID: 123}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	customField, _, err := client.ProductCustomFields.Show(context.Background(), 12, 3)
	assert.Nil(t, err)
	assert.Equal(t, expected, customField)
}
