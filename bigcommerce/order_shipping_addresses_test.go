package bigcommerce

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderShippingAddressService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/12/shipping_addresses/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"page": "1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{ "id": 123 }]`)
	})

	expected := []OrderShippingAddress{
		{ID: 123},
	}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	params := &OrderShippingAddressListParams{
		Page: 1,
	}
	orderShippingAddresses, _, err := client.OrderShippingAddresses.List(context.Background(), 12, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, orderShippingAddresses)
}

func TestOrderShippingAddressService_Count(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/123/shipping_addresses/count", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"limit": "10"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{ "count": 12 }`)
	})

	expected := 12

	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	params := &OrderShippingAddressListParams{
		Limit: 10,
	}
	orderID := 123
	count, _, err := client.OrderShippingAddresses.Count(context.Background(), orderID, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, count)
}

func TestOrderShippingAddressService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/12/shipping_addresses/3", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{ "id": 123 }`)
	})

	expected := &OrderShippingAddress{ID: 123}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	orderShippingAddress, _, err := client.OrderShippingAddresses.Show(context.Background(), 12, 3)
	assert.Nil(t, err)
	assert.Equal(t, expected, orderShippingAddress)
}
