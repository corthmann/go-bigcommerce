package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"context"

	goquery "github.com/google/go-querystring/query"
)

const (
	userAgent  = "go-bigcommerce"
	methodGET  = "GET"
	methodPOST = "POST"
	methodPUT  = "PUT"
)

// Client is a Bigcommerce client for making Bigcommerce API requests.
type Client struct {
	httpClient *http.Client
	// Bigcommerce API Services
	Orders                 *OrderService
	OrderShippingAddresses *OrderShippingAddressService
	OrderStatuses          *OrderStatusService
	Products               *ProductService
	ProductCustomFields    *ProductCustomFieldService
}

// ClientConfig is used to configure the api connection.
type ClientConfig struct {
	Endpoint string `json:"endpoint,omitempty"`
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, config *ClientConfig) *Client {
	return &Client{
		Orders:                 newOrderService(config, httpClient),
		OrderShippingAddresses: newOrderShippingAddressService(config, httpClient),
		OrderStatuses:          newOrderStatusService(config, httpClient),
		Products:               newProductService(config, httpClient),
		ProductCustomFields:    newProductCustomFieldService(config, httpClient),
	}
}

// performGET creates a new context aware HTTP GET request and returns the response.
func performGET(ctx context.Context, httpClient *http.Client, config *ClientConfig, path string, queryParams interface{}, successV, failureV interface{}) (*http.Response, error) {
	return performRequest(ctx, httpClient, config, methodGET, path, queryParams, nil, successV, failureV)
}

// performPOST creates a new context aware HTTP POST request and returns the response.
func performPOST(ctx context.Context, httpClient *http.Client, config *ClientConfig, path string, queryParams interface{}, body interface{}, successV, failureV interface{}) (*http.Response, error) {
	return performRequest(ctx, httpClient, config, methodPOST, path, queryParams, body, successV, failureV)
}

// performPUT creates a new context aware HTTP PUT request and returns the response.
func performPUT(ctx context.Context, httpClient *http.Client, config *ClientConfig, path string, queryParams interface{}, body interface{}, successV, failureV interface{}) (*http.Response, error) {
	return performRequest(ctx, httpClient, config, methodPUT, path, queryParams, body, successV, failureV)
}

// performRequest creates a new context aware HTTP request and returns the response.
func performRequest(ctx context.Context, httpClient *http.Client, config *ClientConfig, method string, path string, queryParams interface{}, body interface{}, successV, failureV interface{}) (*http.Response, error) {
	// Marshal payload
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	// Generate Request Url with query params
	queryValues, err := goquery.Values(queryParams)
	if err != nil {
		return nil, err
	}
	queryString := queryValues.Encode()
	url := fmt.Sprintf("%v/api/v2/%v", config.Endpoint, path)
	if queryString != "" {
		url = strings.Join([]string{url, queryString}, "?")
	}
	// Create Request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	// Set Headers
	req.Header.Add("Accept", "application/json; charset=utf-8")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", userAgent)
	req.SetBasicAuth(config.UserName, config.Password)
	// Perform request
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	// when err is nil, resp contains a non-nil resp.Body which must be closed
	defer response.Body.Close()
	if successV != nil || failureV != nil {
		err = decodeResponseJSON(response, successV, failureV)
	}
	return response, err
}

// decodeResponse decodes response Body into the value pointed to by successV
// if the response is a success (2XX) or into the value pointed to by failureV
// otherwise. If the successV or failureV argument to decode into is nil,
// decoding is skipped.
// Caller is responsible for closing the resp.Body.
func decodeResponseJSON(resp *http.Response, successV, failureV interface{}) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		if successV != nil {
			return decodeResponseBodyJSON(resp, successV)
		}
	} else {
		if failureV != nil {
			if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
				return decodeResponseBodyJSON(resp, failureV)
			}
			return fmt.Errorf("bigcommerce: %v", resp.Status)
		}
	}
	return nil
}

// decodeResponseBodyJSON JSON decodes a Response Body into the value pointed
// to by v.
// Caller must provide a non-nil v and close the resp.Body.
func decodeResponseBodyJSON(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}
