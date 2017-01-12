/*
Package bigcommerce provides an api client for communicating with Bigcommerce REST APIs V2.
The official API documentation can be found on: https://developer.bigcommerce.com/api/v2/

Configure and initialize the client:

  config := &bigcommerce.ClientConfig{
    Endpoint: "https://example.bigcommerce.com",
    UserName: "go-bigcommerce",
    Password: "12345"}
  client := bigcommerce.NewClient(http.DefaultClient, config)

Products

Request a list of products with ID >= 2

	products, resp, err := client.Products.List(context.Background(), &bigcommerce.ProductListParams{
    MinID: 2,
  })

ProductCustomFields

Request a list of ProductCustomFields for products with ID >= 2

  customFields, resp, err := client.ProductCustomFields.List(context.Background(), 2, &bigcommerce.ProductCustomFieldListParams{
    Page: 1,
  })

Orders

Request a list of orders with ID >= 2

  orders, resp, err := client.Orders.List(context.Background(), &bigcommerce.OrderListParams{
    MinID: 2,
  })

OrderShippingAddresses

Request a list of order shipping addresses for Order with ID = 12

  orderShippingAddresses, resp, err := client.OrderShippingAddresses.List(context.Background(), 12, &bigcommerce.OrderShippingAddressesListParams{})

OrderStatuses

Request a list of order statuses

  orderStatuses, resp, err := client.OrderStatuses.List(context.Background(), &bigcommerce.OrderStatusListParams{})

*/
package bigcommerce
