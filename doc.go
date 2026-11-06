// Package asin is the official Go SDK for the Asin API (https://www.asin.dev).
//
// Track Amazon prices and products with a scheduled watchlist, price history, offers and alerts.
//
// Create one client and reuse it:
//
//	client, err := asin.NewClient("", nil) // reads ASIN_API_KEY
//	response, err := client.SearchAmazon(ctx, ...)
//
// Every documented endpoint has a typed method; Client.Request reaches any other path.
package asin
