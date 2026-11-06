package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asindotdev/asin-go"
)

func main() {
	client, err := asin.NewClient("", nil) // reads ASIN_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.Request(ctx, "GET", "/account", nil)
	if apiError, ok := asin.IsAPIError(err); ok {
		log.Fatalf("HTTP %d: %s", apiError.StatusCode, apiError.Message)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response.Body))
}
