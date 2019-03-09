// Package types contains the data structures for unmarshaling data
// returned when querying the commercetools REST API
package types

// ProductProjectionsResults for unmarshaling data from hitting the
// ${baseURL}/project-projections/search?text.en=${terms} endpoint
type ProductProjectionsResults struct {
	Limit   int      `json:"limit"`
	Offset  int      `json:"offset"`
	Count   int      `json:"count"`
	Total   int      `json:"total"`
	Results []Result `json:"results"`
}

// Result contains a single search result from hitting the
// ${baseURL}/product-projections/search endpoint
type Result struct {
	LastModifiedAt     string            `json:"lastModifiedAt"`
	CreatedAt          string            `json:"createdAt"`
	Published          bool              `json:"published"`
	HasStagedChanges   bool              `json:"hasStagedChanges"`
	SearchKeywords     map[string]string `json:"searchKeywords"`
	MasterVariant      Variant           `josn:"masterVariant"`
	Variants           []Variant         `json:"variants"`
	Slug               map[string]string `json:"slug"`
	CategoryOrderHints map[string]string `json:"categoryOrderHints"`
	Categories         []baseType        `json:"categories"`
	Name               map[string]string `json:"name"`
	Version            int               `json:"version"`
	ID                 string            `json:"id"`
	ProductType        baseType          `json:"productType"`
	TaxCategory        baseType          `json:"taxCategory"`
}

// Variant unmarshals the "variant" data structure
type Variant struct {
	Attributes []VariantAttribute `json:"attributes"`
	Prices     []Price            `json:"prices"`
	SKU        string             `json:"sku"`
	ID         int                `json:"id"`
}

// VariantAttribute unmarshals the "attributes" of a "variant"
type VariantAttribute struct {
	Name string `json:"name"`
	// Because the response for "value" can either be
	// a single string, or another nested layer of JSON
	// e.g. "value" : {"key":"spektre", "label": "Spektre"}
	// Because they have the same key in the same layer of JSON,
	// impossible to differentiate while unmarshaling
	// So have to unmarshal into interface{} and coerce type later
	// e.g. if Value.(*string) {}
	// Unmarshaling a nested set of json key/values as an interface{} returns:
	// e.g. {Name:color Value:map[key:oliv label:map[de:oliv en:oliv it:oliva]]}
	Value interface{} `json:"value"`
}

// Price unmarshals the "price" data structure nested within "variants"
type Price struct {
	Value         PriceValue `json:"value"`
	ID            string     `json:"id"`
	Country       string     `json:"country"`
	CustomerGroup baseType   `json:"customerGroup"`
	Channel       baseType   `json:"channel"`
}

// PriceValue unmarshals the "value" data structure nested within "price"
type PriceValue struct {
	Type           string `json:"centPrecision"`
	CurrencyCode   string `json:"currencyCode"`
	CentAmount     int64  `json:"centAmount"`
	FractionDigits int    `json:"fractionDigits"`
}

// Channel unmarshals the "channel" data structure
type Channel struct {
	TypeID string `json:"typeId"` //
	ID     string `json:"id"`
}
