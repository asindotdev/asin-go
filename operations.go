// Code generated from the Asin OpenAPI document. DO NOT EDIT.

package asin

import "context"

// CheckProduct — Check the price now. Asks for a check right away (1 credit), at most once per 5 minutes; earlier calls return the product unchanged. A paused product has to be resumed first. The check runs in the background; with wait=true the answer waits for it. Read last_e…
//
// POST /products/{id}/check
func (c *Client) CheckProduct(ctx context.Context, params *CheckProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["check_product"], params.params(), options)
}

// DeleteProduct — Stop tracking a product. Removes the product and its price history.
//
// DELETE /products/{id}
func (c *Client) DeleteProduct(ctx context.Context, params *DeleteProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["delete_product"], params.params(), options)
}

// FetchOffers — Read every offer from Amazon. Reads the offers on Amazon for 1 credit, refunded if it doesn't answer. A read made in the last 3 hours, by anyone, is returned instead, for free (charged: false). The answer is the same as GET /offers.
//
// POST /offers
func (c *Client) FetchOffers(ctx context.Context, params *FetchOffersParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fetch_offers"], params.params(), options)
}

// GetAccount — Plan and remaining credits. Your plan, the credits left (monthly allowance and packs), how many products you track and the credits they use per month.
//
// GET /account
func (c *Client) GetAccount(ctx context.Context, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["get_account"], nil, options)
}

// GetAsinHistory — Price history of any ASIN. Every price Asin read for an ASIN on a marketplace over the period, whoever tracks it, oldest first, with the lowest, highest and average of those prices. A history of more than 5000 points is thinned evenly, keeping its first and last poi…
//
// GET /history
func (c *Client) GetAsinHistory(ctx context.Context, params *GetAsinHistoryParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["get_asin_history"], params.params(), options)
}

// GetOffers — Offers already known. Who sells the product and at what price, as last read, cheapest first, with Amazon's own price and the cheapest new and used offers (used is only known for Amazon's own outlet). Free: this never reads Amazon or spends a credit. When nobody…
//
// GET /offers
func (c *Client) GetOffers(ctx context.Context, params *GetOffersParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["get_offers"], params.params(), options)
}

// GetProduct — Get a product and its price history. The product's latest state and every price read of the period, oldest first. The history covers every read of this ASIN on this marketplace, including those made before you added it, so a product added today can come with months of prices;…
//
// GET /products/{id}
func (c *Client) GetProduct(ctx context.Context, params *GetProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["get_product"], params.params(), options)
}

// ListProducts — List tracked products. Your watchlist, newest first, with each product's latest state and its recent prices (trend). Filter with asin (and marketplace) to know if a product is tracked, or with filter to get the same lists as the app's tabs.
//
// GET /products
func (c *Client) ListProducts(ctx context.Context, params *ListProductsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["list_products"], params.params(), options)
}

// SearchAmazon — Search Amazon. Amazon search results (up to 60, sponsored included), to find the ASIN of a product. Free, limited to 30 searches per hour.
//
// GET /search
func (c *Client) SearchAmazon(ctx context.Context, params *SearchAmazonParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["search_amazon"], params.params(), options)
}

// TrackProduct — Track a product. Adds an ASIN, an Amazon link or a short link (amzn.to, a.co…) to your watchlist. The first check runs in the background within seconds (checking is true until then): poll GET /products/{id}, or send wait: true to get the answer once it's d…
//
// POST /products
func (c *Client) TrackProduct(ctx context.Context, params *TrackProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["track_product"], params.params(), options)
}

// UpdateProduct — Change alerts and checks. Target price, drop alerts, check interval, or pause. Send only what changes.
//
// PATCH /products/{id}
func (c *Client) UpdateProduct(ctx context.Context, params *UpdateProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["update_product"], params.params(), options)
}

// endpoints lists every documented operation by its snake_case name.
var endpoints = map[string]endpoint{
	"check_product":    {path: "/products/{id}/check", methods: []string{"POST"}, json: true, required: []string{"id"}},
	"delete_product":   {path: "/products/{id}", methods: []string{"DELETE"}, json: true, required: []string{"id"}},
	"fetch_offers":     {path: "/offers", methods: []string{"POST"}, json: true, required: []string{"asin"}},
	"get_account":      {path: "/account", methods: []string{"GET"}, json: true, required: []string{}},
	"get_asin_history": {path: "/history", methods: []string{"GET"}, json: true, required: []string{"asin"}},
	"get_offers":       {path: "/offers", methods: []string{"GET"}, json: true, required: []string{"asin"}},
	"get_product":      {path: "/products/{id}", methods: []string{"GET"}, json: true, required: []string{"id"}},
	"list_products":    {path: "/products", methods: []string{"GET"}, json: true, required: []string{}},
	"search_amazon":    {path: "/search", methods: []string{"GET"}, json: true, required: []string{"query"}},
	"track_product":    {path: "/products", methods: []string{"POST"}, json: true, required: []string{"product"}},
	"update_product":   {path: "/products/{id}", methods: []string{"PATCH"}, json: true, required: []string{"id"}},
}
