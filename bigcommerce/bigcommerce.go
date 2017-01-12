package bigcommerce

import (
	"encoding/json"
	"net/http"

	"github.com/dghubble/sling"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

const (
	userAgent = "go-bigcommerce"
)

// Client is a Bigcommerce client for making Bigcommerce API requests.
type Client struct {
	sling      *sling.Sling
	httpClient *http.Client
	// Bigcommerce API Services
	Orders        *OrderService
	OrderStatuses *OrderStatusService
	Products      *ProductService
}

// ClientConfig is used to configure the api connection.
type ClientConfig struct {
	Endpoint string `json:"endpoint,omitempty"`
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, config *ClientConfig) *Client {
	base := sling.New().Client(httpClient).SetBasicAuth(config.UserName, config.Password).Set("Accept", "application/json; charset=utf-8").Set("Content-Type", "application/json").Base(config.Endpoint + "/api/v2/")
	return &Client{
		sling:         base,
		Orders:        newOrderService(base.New()),
		OrderStatuses: newOrderStatusService(base.New()),
		Products:      newProductService(base.New(), httpClient),
	}
}

// performRequest creates a new context aware HTTP request and returns the response.
func performRequest(ctx context.Context, s *sling.Sling, httpClient *http.Client, successV, failureV interface{}) (*http.Response, error) {
	req, err := s.Request()
	if err != nil {
		return nil, err
	}
	resp, err := ctxhttp.Do(ctx, httpClient, req)
	if err != nil {
		return resp, err
	}
	// when err is nil, resp contains a non-nil resp.Body which must be closed
	defer resp.Body.Close()
	if successV != nil || failureV != nil {
		err = decodeResponseJSON(resp, successV, failureV)
	}
	return resp, err
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
			return decodeResponseBodyJSON(resp, failureV)
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
