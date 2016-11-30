package bigcommerce

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	userAgent = "go-bigcommerce"
)

// Client is a Twitter client for making Twitter API requests.
type Client struct {
	sling *sling.Sling
	// Bigcommerce API Services
	Products *ProductService
}

type ClientConfig struct {
	Endpoint string `json:"storeHash,omitempty"`
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

/// Incomplete stuff.... / inspiration
//
// type Configuration struct {
// 	UserName      string            `json:"userName,omitempty"`
// 	Password      string            `json:"password,omitempty"`
// 	APIKeyPrefix  map[string]string `json:"APIKeyPrefix,omitempty"`
// 	APIKey        map[string]string `json:"APIKey,omitempty"`
// 	Debug         bool              `json:"debug,omitempty"`
// 	DebugFile     string            `json:"debugFile,omitempty"`
// 	OAuthToken    string            `json:"oAuthToken,omitempty"`
// 	BasePath      string            `json:"basePath,omitempty"`
// 	Host          string            `json:"host,omitempty"`
// 	Scheme        string            `json:"scheme,omitempty"`
// 	AccessToken   string            `json:"accessToken,omitempty"`
// 	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
// 	UserAgent     string            `json:"userAgent,omitempty"`
// 	APIClient     *APIClient
// 	Transport     *http.Transport
// 	Timeout       *time.Duration    `json:"timeout,omitempty"`
// }
//
// func NewConfiguration() *Configuration {
// 	cfg := &Configuration{
// 		BasePath:      "https://api.bigcommerce.com/stores/{{store_id}}/v3",
// 		DefaultHeader: make(map[string]string),
// 		APIKey:        make(map[string]string),
// 		APIKeyPrefix:  make(map[string]string),
// 		UserAgent:     "Swagger-Codegen/1.0.0/go",
// 		APIClient:     &APIClient{},
// 	}
//
// 	cfg.APIClient.config = cfg
// 	return cfg
// }
//
// func (c *Configuration) GetBasicAuthEncodedString() string {
// 	return base64.StdEncoding.EncodeToString([]byte(c.UserName + ":" + c.Password))
// }
