# go-bigcommerce
Go Bigcommerce REST API

Installation
-------------
```
go get github.com/corthmann/go-bigcommerce
```

Usage
-------------
```
import { "bigcommerce" }
config := &bigcommerce.ClientConfig{
  Endpoint: "https://example.bigcommerce.com",
  UserName: "go-bigcommerce",
  Password: "12345"}
client := bigcommerce.NewClient(http.DefaultClient, config)

products, resp, err := client.Products.List(&bigcommerce.ProductListParams{
  Sku: "Product-123",
  })
```
