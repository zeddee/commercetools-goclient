package types

// TaxCategory unmarshals "taxCategory"
// Also has a "draft" form
type TaxCategory struct {
	ID             string    `json:"id"`
	Key            string    `json:"key"`
	Version        int       `json:"version"`
	CreatedAt      string    `json:"createdAd"`
	LastModifiedAt string    `json:"lastModifiedAt"`
	Name           string    `json:"name"`
	Description    string    `json:"description"` // optional
	Rates          []TaxRate `json:"rates"`
}

// Rate describes a percentage value, as in a 0.5 tax Rate applied to a product price is a 50% tax
// Should only allow values [0.. 1]
type Rate float32

// TaxRate unmarshals "taxRate"
// Also has a "draft" form
type TaxRate struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Amount          Rate      `json:"amount"`
	IncludedInPrice bool      `json:"includedInPrice"`
	Country         string    `json:"country"`  // A two-character country code as per ↗ ISO 3166-1 alpha-2 .
	State           string    `json:"state"`    // optional
	SubRates        []SubRate `json:"subRates"` // For countries (e.g. the US) where the total tax is a combination of multiple taxes (e.g. state and local taxes).
}

// SubRate unmarshals "subRate"
type SubRate struct {
	Name   string `json:"name"`
	Amount Rate   `json:"amount"`
}
