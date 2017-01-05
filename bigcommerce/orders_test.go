package bigcommerce

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"customer_id": "0"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{ "id": 123 }]`)
	})

	expected := &Orders{
		{ID: 123},
	}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	customerID := uint32(0)
	params := &OrderListParams{
		CustomerID: &customerID,
	}
	orders, _, err := client.Orders.List(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, orders)
}

func TestOrderService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{ "id": 123 }`)
	})

	expected := &Order{ID: 123}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	order, _, err := client.Orders.Show(123)
	assert.Nil(t, err)
	assert.Equal(t, expected, order)
}
