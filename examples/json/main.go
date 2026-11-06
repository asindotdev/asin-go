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
	response, err := client.SearchAmazon(ctx, &asin.SearchAmazonParams{Query: "wireless earbuds"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response.Body))
}
