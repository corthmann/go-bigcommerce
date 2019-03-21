package bigcommerce

import (
	"context"
	"encoding/json"
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

func TestOrderService_ListWithError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"customer_id": "0"}, r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, BadRequestJSON)
	})

	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	customerID := 0
	params := &OrderListParams{
		CustomerID: &customerID,
	}
	orders, _, err := client.Orders.List(context.Background(), params)
	assert.EqualError(t, err, BadRequestErrorMessage)
	assert.True(t, len(orders) == 0)
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

func TestOrderService_CountWithError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/count", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"customer_id": "0"}, r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, BadRequestJSON)
	})

	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	customerID := 0
	params := &OrderListParams{
		CustomerID: &customerID,
	}
	_, _, err := client.Orders.Count(context.Background(), params)
	assert.EqualError(t, err, BadRequestErrorMessage)
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

func TestOrderService_ShowWithError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, BadRequestJSON)
	})
	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	_, _, err := client.Orders.Show(context.Background(), 123)
	assert.EqualError(t, err, BadRequestErrorMessage)
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
				FirstName:      "Tester",
				LastName:       "Test",
				Street1:        "Test Street 1",
				Street2:        "",
				Zip:            "1234",
				City:           "Test City",
				Country:        "Denmark",
				CountryIso2:    "dk",
				Phone:          "12345678",
				Email:          "example@example.com",
				ShippingMethod: "2-Day",
			},
		},
		CustomerMessage: "This is a test.",
		StaffNotes:      "Yeah... I'm not buying it.",
	}
	order, _, err := client.Orders.New(context.Background(), body)
	assert.Nil(t, err)
	assert.Equal(t, expected, order)
}

func TestOrderService_NewWithError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, BadRequestJSON)
	})

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
				FirstName:      "Tester",
				LastName:       "Test",
				Street1:        "Test Street 1",
				Street2:        "",
				Zip:            "1234",
				City:           "Test City",
				Country:        "Denmark",
				CountryIso2:    "dk",
				Phone:          "12345678",
				Email:          "example@example.com",
				ShippingMethod: "2-Day",
			},
		},
		CustomerMessage: "This is a test.",
		StaffNotes:      "Yeah... I'm not buying it.",
	}
	_, _, err := client.Orders.New(context.Background(), body)
	assert.EqualError(t, err, BadRequestErrorMessage)
}

func TestOrderService_NewReply(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
    "id": 100,
    "customer_id": 10,
    "date_created": "Wed, 14 Nov 2012 19:26:23 +0000",
    "date_modified": "Wed, 14 Nov 2012 19:26:23 +0000",
    "date_shipped": "",
    "status_id": 11,
    "status": "Awaiting Fulfillment",
    "subtotal_ex_tax": "79.0000",
    "subtotal_inc_tax": "79.0000",
    "subtotal_tax": "0.0000",
    "base_shipping_cost": "0.0000",
    "shipping_cost_ex_tax": "0.0000",
    "shipping_cost_inc_tax": "0.0000",
    "shipping_cost_tax": "0.0000",
    "shipping_cost_tax_class_id": 2,
    "base_handling_cost": "0.0000",
    "handling_cost_ex_tax": "0.0000",
    "handling_cost_inc_tax": "0.0000",
    "handling_cost_tax": "0.0000",
    "handling_cost_tax_class_id": 2,
    "base_wrapping_cost": "0.0000",
    "wrapping_cost_ex_tax": "0.0000",
    "wrapping_cost_inc_tax": "0.0000",
    "wrapping_cost_tax": "0.0000",
    "wrapping_cost_tax_class_id": 3,
    "total_ex_tax": "79.0000",
    "total_inc_tax": "79.0000",
    "total_tax": "0.0000",
    "items_total": 1,
    "items_shipped": 0,
    "payment_method": "cash",
    "payment_provider_id": null,
    "payment_status": "",
    "refunded_amount": "0.0000",
    "order_is_digital": false,
    "store_credit_amount": "0.0000",
    "gift_certificate_amount": "0.0000",
    "ip_address": "50.58.18.2",
    "geoip_country": "",
    "geoip_country_iso2": "",
    "currency_id": 1,
    "currency_code": "USD",
    "currency_exchange_rate": "1.0000000000",
    "default_currency_id": 1,
    "default_currency_code": "USD",
    "staff_notes": "",
    "customer_message": "",
    "discount_amount": "0.0000",
    "coupon_discount": "0.0000",
    "shipping_address_count": 1,
    "is_deleted": false,
    "billing_address": {
      "first_name": "Trisha",
      "last_name": "McLaughlin",
      "company": "",
      "street_1": "12345 W Anderson Ln",
      "street_2": "",
      "city": "Austin",
      "state": "Texas",
      "zip": "78757",
      "country": "United States",
      "country_iso2": "US",
      "phone": "",
      "email": "elsie@example.com"
    },
    "products": {
      "url": "https://store-bwvr466.mybigcommerce.com/api/v2/orders/100/products.json",
      "resource": "/orders/100/products"
    },
    "shipping_addresses": {
      "url": "https://store-bwvr466.mybigcommerce.com/api/v2/orders/100/shippingaddresses.json",
      "resource": "/orders/100/shippingaddresses"
    },
    "coupons": {
      "url": "https://store-bwvr466.mybigcommerce.com/api/v2/orders/100/coupons.json",
      "resource": "/orders/100/coupons"
    }
  }`)
	})

	want := `{"id":100,"customer_id":10,"date_created":"Wed, 14 Nov 2012 19:26:23 +0000","date_modified":"Wed, 14 Nov 2012 19:26:23 +0000","date_shipped":"","status_id":11,"status":"Awaiting Fulfillment","handling_cost_ex_tax":"0","handling_cost_inc_tax":"0","handling_cost_tax":"0","shipping_cost_ex_tax":"0","shipping_cost_inc_tax":"0","shipping_cost_tax":"0","subtotal_ex_tax":"79","subtotal_inc_tax":"79","subtotal_tax":"0","total_ex_tax":"79","total_inc_tax":"79","total_tax":"0","base_shipping_cost":"0","items_total":1,"payment_method":"cash","payment_status":"","ip_address":"50.58.18.2","currency_id":1,"currency_code":"USD","staff_notes":"","customer_message":"","discount_amount":"0.0000","counpon_discount":"","shipping_address_count":1,"billing_address":{"first_name":"Trisha","last_name":"McLaughlin","company":"","street_1":"12345 W Anderson Ln","street_2":"","city":"Austin","state":"Texas","zip":"78757","country":"United States","country_iso2":"US","phone":"","email":"elsie@example.com"}}`
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
	got, err := json.Marshal(order)
	assert.Nil(t, err)
	assert.Equal(t, want, string(got))
	assert.Nil(t, order.DateShipped.Time())
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
	statusID := 1
	params := &OrderEditParams{
		StatusID: &statusID,
	}
	order, _, err := client.Orders.Edit(context.Background(), 123, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, order)
}

func TestOrderService_EditWithError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/v2/orders/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadRequestJSON)
	})

	client := NewClient(httpClient, &ClientConfig{
		Endpoint: "https://example.com",
		UserName: "go-bigcommerce",
		Password: "12345"})
	statusID := 1
	params := &OrderEditParams{
		StatusID: &statusID,
	}
	_, _, err := client.Orders.Edit(context.Background(), 123, params)
	assert.EqualError(t, err, BadRequestErrorMessage)
}
