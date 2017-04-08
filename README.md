# go-bigcommerce
Go Bigcommerce REST API

[![Build Status](https://travis-ci.org/corthmann/go-bigcommerce.svg?branch=master)](https://travis-ci.org/corthmann/go-bigcommerce)

Installation
-------------
```
go get github.com/corthmann/go-bigcommerce/bigcommerce
```

Usage
-------------
See [GoDoc](https://godoc.org/github.com/corthmann/go-bigcommerce/bigcommerce) for API and examples.

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
