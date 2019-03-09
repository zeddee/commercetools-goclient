package types

// baseType is a general data structure that commercetools uses
// to describe certain properties attached to larger concepts e.g. products
// It typically contains a "typeId" and a "id" field.
// ProductType unmarshals a productType
type baseType struct {
	TypeID string `json:"typeId"`
	ID     string `json:"id"`
}
