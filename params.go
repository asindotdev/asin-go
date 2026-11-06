// Code generated from the Asin OpenAPI document. DO NOT EDIT.

package asin

// CheckProductParams holds the parameters of CheckProduct (POST /products/{id}/check).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type CheckProductParams struct {
	// product_id returned when the product was added
	ID string `json:"id" api:"id,required"`
	// Answer once the check is done (up to 45 seconds) Default: false.
	Wait *bool `json:"wait" api:"wait"`
}

func (p *CheckProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 2)
	if p.ID != "" {
		values = append(values, param{"id", p.ID})
	}
	if p.Wait != nil {
		values = append(values, param{"wait", *p.Wait})
	}
	return values
}

// DeleteProductParams holds the parameters of DeleteProduct (DELETE /products/{id}).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type DeleteProductParams struct {
	// product_id returned when the product was added
	ID string `json:"id" api:"id,required"`
}

func (p *DeleteProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 1)
	if p.ID != "" {
		values = append(values, param{"id", p.ID})
	}
	return values
}

// FetchOffersParams holds the parameters of FetchOffers (POST /offers).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type FetchOffersParams struct {
	// 10-character ASIN
	Asin string `json:"asin" api:"asin,required"`
	// Amazon store, e.g. com for amazon.com Allowed values: com, co.uk, de, fr, it, es, ca, co.jp, com.br, com.mx, in. Default: "com".
	Marketplace *string `json:"marketplace" api:"marketplace"`
	// Read Amazon again unless the last read is under 5 minutes old (1 credit) Default: false.
	Refresh *bool `json:"refresh" api:"refresh"`
}

func (p *FetchOffersParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 3)
	if p.Asin != "" {
		values = append(values, param{"asin", p.Asin})
	}
	if p.Marketplace != nil {
		values = append(values, param{"marketplace", *p.Marketplace})
	}
	if p.Refresh != nil {
		values = append(values, param{"refresh", *p.Refresh})
	}
	return values
}

// GetAsinHistoryParams holds the parameters of GetAsinHistory (GET /history).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type GetAsinHistoryParams struct {
	// 10-character ASIN
	Asin string `json:"asin" api:"asin,required"`
	// Amazon store, e.g. com for amazon.com Allowed values: com, co.uk, de, fr, it, es, ca, co.jp, com.br, com.mx, in. Default: "com".
	Marketplace *string `json:"marketplace" api:"marketplace"`
	// History period in days Default: 90.
	Days *int64 `json:"days" api:"days"`
}

func (p *GetAsinHistoryParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 3)
	if p.Asin != "" {
		values = append(values, param{"asin", p.Asin})
	}
	if p.Marketplace != nil {
		values = append(values, param{"marketplace", *p.Marketplace})
	}
	if p.Days != nil {
		values = append(values, param{"days", *p.Days})
	}
	return values
}

// GetOffersParams holds the parameters of GetOffers (GET /offers).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type GetOffersParams struct {
	// 10-character ASIN
	Asin string `json:"asin" api:"asin,required"`
	// Amazon store, e.g. com for amazon.com Allowed values: com, co.uk, de, fr, it, es, ca, co.jp, com.br, com.mx, in. Default: "com".
	Marketplace *string `json:"marketplace" api:"marketplace"`
}

func (p *GetOffersParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 2)
	if p.Asin != "" {
		values = append(values, param{"asin", p.Asin})
	}
	if p.Marketplace != nil {
		values = append(values, param{"marketplace", *p.Marketplace})
	}
	return values
}

// GetProductParams holds the parameters of GetProduct (GET /products/{id}).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type GetProductParams struct {
	// product_id returned when the product was added
	ID string `json:"id" api:"id,required"`
	// History period in days Default: 90.
	Days *int64 `json:"days" api:"days"`
}

func (p *GetProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 2)
	if p.ID != "" {
		values = append(values, param{"id", p.ID})
	}
	if p.Days != nil {
		values = append(values, param{"days", *p.Days})
	}
	return values
}

// ListProductsParams holds the parameters of ListProducts (GET /products).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type ListProductsParams struct {
	// Only this ASIN
	Asin *string `json:"asin" api:"asin"`
	// Only this Amazon store Allowed values: com, co.uk, de, fr, it, es, ca, co.jp, com.br, com.mx, in.
	Marketplace *string `json:"marketplace" api:"marketplace"`
	// Only the products of one of the app's watchlist filters. all: Every tracked product (default). drops: Cheaper now than when the product was added. target: At or below its target price. out_of_stock: No offer at the last check. attention: The last check didn't go through (last_error is set): no cred…
	Filter *string `json:"filter" api:"filter"`
	// Page number Default: 1.
	Page *int64 `json:"page" api:"page"`
	// Products per page Default: 20.
	Limit *int64 `json:"limit" api:"limit"`
}

func (p *ListProductsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Asin != nil {
		values = append(values, param{"asin", *p.Asin})
	}
	if p.Marketplace != nil {
		values = append(values, param{"marketplace", *p.Marketplace})
	}
	if p.Filter != nil {
		values = append(values, param{"filter", *p.Filter})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Limit != nil {
		values = append(values, param{"limit", *p.Limit})
	}
	return values
}

// SearchAmazonParams holds the parameters of SearchAmazon (GET /search).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type SearchAmazonParams struct {
	// Keywords
	Query string `json:"query" api:"query,required"`
	// Amazon store, e.g. com for amazon.com Allowed values: com, co.uk, de, fr, it, es, ca, co.jp, com.br, com.mx, in. Default: "com".
	Marketplace *string `json:"marketplace" api:"marketplace"`
}

func (p *SearchAmazonParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 2)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Marketplace != nil {
		values = append(values, param{"marketplace", *p.Marketplace})
	}
	return values
}

// TrackProductParams holds the parameters of TrackProduct (POST /products).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type TrackProductParams struct {
	// ASIN (B09DFCB66S) or Amazon product link
	Product string `json:"product" api:"product,required"`
	// Amazon store, e.g. com for amazon.com. A product link sets it on its own. Allowed values: com, co.uk, de, fr, it, es, ca, co.jp, com.br, com.mx, in. Default: "com".
	Marketplace *string `json:"marketplace" api:"marketplace"`
	// Email the owner when the price goes at or below this amount. null removes it.
	TargetPrice *float64 `json:"target_price" api:"target_price"`
	// Minutes between two checks (1440 = daily). Checks faster than daily need a plan or a credit pack: every 60 minutes with a pack or Starter, 30 with Pro, 5 with Business. Default: 1440.
	IntervalMinutes *int64 `json:"interval_minutes" api:"interval_minutes"`
	// Answer once the first check is done (up to 45 seconds) Default: false.
	Wait *bool `json:"wait" api:"wait"`
}

func (p *TrackProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Product != "" {
		values = append(values, param{"product", p.Product})
	}
	if p.Marketplace != nil {
		values = append(values, param{"marketplace", *p.Marketplace})
	}
	if p.TargetPrice != nil {
		values = append(values, param{"target_price", *p.TargetPrice})
	}
	if p.IntervalMinutes != nil {
		values = append(values, param{"interval_minutes", *p.IntervalMinutes})
	}
	if p.Wait != nil {
		values = append(values, param{"wait", *p.Wait})
	}
	return values
}

// UpdateProductParams holds the parameters of UpdateProduct (PATCH /products/{id}).
// Required fields are plain values; optional fields are pointers (use asin.String, asin.Float, asin.Bool) or nil slices and maps.
type UpdateProductParams struct {
	// product_id returned when the product was added
	ID string `json:"id" api:"id,required"`
	// Email the owner when the price goes at or below this amount. null removes it.
	TargetPrice *float64 `json:"target_price" api:"target_price"`
	// Minutes between two checks (1440 = daily). Checks faster than daily need a plan or a credit pack: every 60 minutes with a pack or Starter, 30 with Pro, 5 with Business. Default: 1440.
	IntervalMinutes *int64 `json:"interval_minutes" api:"interval_minutes"`
	// Email when the price drops or the product is back in stock
	AlertOnDrop *bool `json:"alert_on_drop" api:"alert_on_drop"`
	// Stop checking (and charging) for now; false resumes
	Paused *bool `json:"paused" api:"paused"`
}

func (p *UpdateProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.ID != "" {
		values = append(values, param{"id", p.ID})
	}
	if p.TargetPrice != nil {
		values = append(values, param{"target_price", *p.TargetPrice})
	}
	if p.IntervalMinutes != nil {
		values = append(values, param{"interval_minutes", *p.IntervalMinutes})
	}
	if p.AlertOnDrop != nil {
		values = append(values, param{"alert_on_drop", *p.AlertOnDrop})
	}
	if p.Paused != nil {
		values = append(values, param{"paused", *p.Paused})
	}
	return values
}
