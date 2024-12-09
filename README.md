# Billingo API's go client

Current version: _3.0.14_

Generated from billingo's OpenAPI spec.

# Usage

### Installation

`go get -u github.com/pilab-dev/billingo/v3`

### Create Product:

```go
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.cm/pilab-dev/go-billingo/v3"
)

func main() {
	apiToken := os.Getenv("BILLINGO_API_TOKEN")

	c, err := billingo.New(apiToken)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create client: %v\n", err)
		os.Exit(1)
	}

	// create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create product
	res, err := c.CreateProductWithResponse(ctx, billingo.Product{
		Comment:              new(string),
		Currency:             "HUF",
		Entitlement:          billingo.ToPtr(billingo.EntitlementAAM),
		GeneralLedgerNumber:  new(string),
		GeneralLedgerTaxcode: new(string),
		Id:                   billingo.ToPtr(1),
		Name:                 "Test product",
		NetUnitPrice:         billingo.ToPtr(float32(600.1)),
		Unit:                 "db",
		Vat:                  "27%",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to call the billingo server: %v\n", err)
		os.Exit(1)
	}

	if res.StatusCode() != 201 {
		fmt.Fprintf(os.Stderr, "expected status code 201, got %d\n", res.StatusCode())
		os.Exit(1)
	}

	fmt.Printf("Created product with ID: %d\n", *res.JSON201.Id)
}
```

## TODO

- Provide more examples

Made with ❤️ by Progressive Innovation LAB
