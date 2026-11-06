# Asin Go SDK

The official Go SDK for the [Asin](https://www.asin.dev) Amazon price tracking API.

Track Amazon prices and products with a scheduled watchlist, price history, offers and alerts.

The SDK exposes **11 typed methods**, one per documented endpoint, plus a generic `Request()`. It has no dependencies outside the standard library.

## Requirements

- Go 1.22 or newer

## Installation

```sh
go get github.com/asindotdev/asin-go
```

## Get your API key

[Sign up for Asin](https://www.asin.dev/signup) or sign in, then copy your API key from the [dashboard](https://www.asin.dev/app). `NewClient("", nil)` reads the `ASIN_API_KEY` environment variable.

## Quick start

```go
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
	data, err := response.Map()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
```

Parameter structs mirror the [API reference](https://www.asin.dev/docs). Required fields are plain values; optional fields are pointers — use `asin.String`, `asin.Float`, `asin.Int` and `asin.Bool`.

Decode into your own types with `response.Decode(&value)`.

## HTTP methods

Each method uses the HTTP verb of its endpoint. `GET` and `DELETE` send parameters in the query string; `POST` and `PATCH` send them as a JSON body, except `wait` on `/products/{id}/check`, which stays in the query string. Null values are never sent, so `target_price` cannot be removed through the SDK yet.

## Any endpoint

```go
response, err := client.Request(ctx, "GET", "/account", nil)
```

## Errors

```go
response, err := client.SearchAmazon(ctx, &asin.SearchAmazonParams{Query: "wireless earbuds"})
if apiError, ok := asin.IsAPIError(err); ok {
	log.Printf("HTTP %d: %s", apiError.StatusCode, apiError.Message)
}
```

| Error | When |
| --- | --- |
| `*ConfigurationError` | Missing API key, invalid options or unsupported method. |
| `*SerializationError` | A required parameter is missing or a value can't be sent. |
| `*TransportError` | Network failure, timeout or cancelled context (`errors.Is(err, context.DeadlineExceeded)`). |
| `*APIError` | Non-2xx status (`StatusCode`, `Message`, `Response`). |
| `*ResponseDecodingError` | A success response could not be decoded. |

## Retries

Requests are retried up to 3 times after a network error, a `408`, `429`, `500`, `502`, `503` or `504`. Other statuses, such as `400`, `401`, `402` or `404`, and client-side timeouts are not retried. The delay starts at 1 second and grows 1.5× per retry (capped at 30 seconds); a `Retry-After` header wins.

```go
client, err := asin.NewClient("", &asin.ClientOptions{
	MaxRetries: 5,                      // negative disables retries
	RetryDelay: 500 * time.Millisecond,
})

// or for one request
response, err := client.SearchAmazon(ctx, &asin.SearchAmazonParams{Query: "wireless earbuds"}, asin.WithMaxRetries(0))
```

## Batch requests

`asin.Batch` runs a function per item with at most 10 calls in flight and returns one `BatchResult` per item, in input order. A failing item never stops the others.

```go
inputs := []string{"wireless earbuds", "mechanical keyboard", "usb-c hub"}
results := asin.Batch(ctx, inputs, 5, func(ctx context.Context, value string) (*asin.Response, error) {
	return client.SearchAmazon(ctx, &asin.SearchAmazonParams{Query: value})
})
for _, result := range results {
	if result.Err != nil {
		log.Printf("%s: %v", inputs[result.Index], result.Err)
		continue
	}
	fmt.Println(string(result.Value.Body))
}
```

| Constant | Default | Description |
| --- | --- | --- |
| `DefaultConcurrency` | `10` | Tasks run at the same time by batch |
| `DefaultTimeout` | `120 * time.Second` | Request timeout |
| `MaxRetries` | `3` | Retry attempts after a network error, a 408/429/5xx |
| `RetryBackoffFactor` | `1.5` | Exponential backoff multiplier |
| `DefaultRetryDelay` | `time.Second` | Delay before the first retry |

## Timeouts and custom HTTP clients

```go
client, err := asin.NewClient("", &asin.ClientOptions{
	Timeout:    60 * time.Second, // default 120 s; negative disables
	HTTPClient: &http.Client{Transport: myTransport},
})
```

Every method also honours the deadline and cancellation of the `context.Context` you pass.

## Endpoints

11 methods, generated from the [OpenAPI document](https://www.asin.dev/openapi.json).

| Method | HTTP | Description |
| --- | --- | --- |
| `CheckProduct` | POST `/products/{id}/check` | Asks for a check right away (1 credit), at most once per 5 minutes; earlier calls return the product unchanged. A paused product has to be… |
| `DeleteProduct` | DELETE `/products/{id}` | Removes the product and its price history. |
| `FetchOffers` | POST `/offers` | Reads the offers on Amazon for 1 credit, refunded if it doesn't answer. A read made in the last 3 hours, by anyone, is returned instead, fo… |
| `GetAccount` | GET `/account` | Your plan, the credits left (monthly allowance and packs), how many products you track and the credits they use per month. |
| `GetAsinHistory` | GET `/history` | Every price Asin read for an ASIN on a marketplace over the period, whoever tracks it, oldest first, with the lowest, highest and average o… |
| `GetOffers` | GET `/offers` | Who sells the product and at what price, as last read, cheapest first, with Amazon's own price and the cheapest new and used offers (used i… |
| `GetProduct` | GET `/products/{id}` | The product's latest state and every price read of the period, oldest first. The history covers every read of this ASIN on this marketplace… |
| `ListProducts` | GET `/products` | Your watchlist, newest first, with each product's latest state and its recent prices (trend). Filter with asin (and marketplace) to know if… |
| `SearchAmazon` | GET `/search` | Amazon search results (up to 60, sponsored included), to find the ASIN of a product. Free, limited to 30 searches per hour. |
| `TrackProduct` | POST `/products` | Adds an ASIN, an Amazon link or a short link (amzn.to, a.co…) to your watchlist. The first check runs in the background within seconds (che… |
| `UpdateProduct` | PATCH `/products/{id}` | Target price, drop alerts, check interval, or pause. Send only what changes. |

## All Asin SDKs

| SDK | Package | Install |
| --- | --- | --- |
| Node.js | [`@asindev/sdk`](https://www.npmjs.com/package/@asindev/sdk) | `npm install @asindev/sdk` |
| Python | [`asindotdev`](https://pypi.org/project/asindotdev/) | `pip install asindotdev` |
| Go | [`asin-go`](https://pkg.go.dev/github.com/asindotdev/asin-go) | `go get github.com/asindotdev/asin-go` |
| .NET | [`Asin`](https://www.nuget.org/packages/Asin) | `dotnet add package Asin` |
| Ruby | [`asindotdev`](https://rubygems.org/gems/asindotdev) | `gem install asindotdev` |
| PHP | [`asindotdev/sdk`](https://packagist.org/packages/asindotdev/sdk) | `composer require asindotdev/sdk` |
| Java | [`asin-java`](https://github.com/asindotdev/asin-java) | build from source, see its README |
| CLI | [`@asindev/cli`](https://www.npmjs.com/package/@asindev/cli) | `npm install -g @asindev/cli` |
| MCP server | [`@asindev/mcp`](https://www.npmjs.com/package/@asindev/mcp) | `npx -y @asindev/mcp` |

## License

MIT — see [LICENSE](LICENSE).
