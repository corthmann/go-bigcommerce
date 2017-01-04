package bigcommerce

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	userAgent = "go-bigcommerce"
)

// Client is a Bigcommerce client for making Bigcommerce API requests.
type Client struct {
	sling *sling.Sling
	// Bigcommerce API Services
	Products *ProductService
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
		sling:    base,
		Products: newProductService(base.New()),
	}
}
