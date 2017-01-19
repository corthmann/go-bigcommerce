package bigcommerce

import (
	"context"
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
		fmt.Fprint(w, `[{ "id": 123 }]`)
	})

	expected := []Order{
		{ID: 123},
	}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	customerID := 0
	params := &OrderListParams{
		CustomerID: &customerID,
	}
	orders, _, err := client.Orders.List(context.Background(), params)
	assert.Nil(t, err)
	assert.Equal(t, expected, orders)
}

func TestOrderService_Count(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/count", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"customer_id": "0"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{ "count": 12 }`)
	})

	expected := 12
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	customerID := 0
	params := &OrderListParams{
		CustomerID: &customerID,
	}
	count, _, err := client.Orders.Count(context.Background(), params)
	assert.Nil(t, err)
	assert.Equal(t, expected, count)
}

func TestOrderService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{ "id": 123 }`)
	})

	expected := &Order{ID: 123}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	order, _, err := client.Orders.Show(context.Background(), 123)
	assert.Nil(t, err)
	assert.Equal(t, expected, order)
}

func TestOrderService_New(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{ "id": 123 }`)
	})

	expected := &Order{ID: 123}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	customerID := 10
	statusID := 0
	body := &OrderBody{
		ExternalSource: "test-suite",
		CustomerID:     &customerID,
		StatusID:       &statusID,
		BillingAddress: AddressEntity{
			FirstName:   "Tester",
			LastName:    "Test",
			Street1:     "Test Street 1",
			Street2:     "",
			Zip:         "1234",
			City:        "Test City",
			Country:     "Denmark",
			CountryIso2: "dk",
			Phone:       "12345678",
			Email:       "example@example.com",
		},
		Products: []OrderProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 1},
		},
		ShippingCostIncTax: 0.0,
		ShippingAddresses: AddressEntities{
			{
				FirstName:   "Tester",
				LastName:    "Test",
				Street1:     "Test Street 1",
				Street2:     "",
				Zip:         "1234",
				City:        "Test City",
				Country:     "Denmark",
				CountryIso2: "dk",
				Phone:       "12345678",
				Email:       "example@example.com",
			},
		},
		CustomerMessage: "This is a test.",
		StaffNotes:      "Yeah... I'm not buying it.",
	}
	order, _, err := client.Orders.New(context.Background(), body)
	assert.Nil(t, err)
	assert.Equal(t, expected, order)
}

func TestOrderService_Edit(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{ "id": 123 }`)
	})

	expected := &Order{ID: 123}
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	statusID := uint32(1)
	params := &OrderEditParams{
		StatusID: &statusID,
	}
	order, _, err := client.Orders.Edit(context.Background(), 123, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, order)
}
