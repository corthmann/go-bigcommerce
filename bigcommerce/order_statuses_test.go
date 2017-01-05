package bigcommerce

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderStatusService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/order_statuses/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"limit": "1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{ "id": 123 }]`)
	})

	expected := &OrderStatuses{
		{ID: 123},
	}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	params := &OrderStatusListParams{
		Limit: 1,
	}
	orderStatuses, _, err := client.OrderStatuses.List(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, orderStatuses)
}

func TestOrderStatusService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/order_statuses/1", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{ "id": 1 }`)
	})

	expected := &OrderStatus{ID: 1}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	orderStatus, _, err := client.OrderStatuses.Show(1)
	assert.Nil(t, err)
	assert.Equal(t, expected, orderStatus)
}
