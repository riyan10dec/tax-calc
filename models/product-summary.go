package models

type ProductSummary struct {
	PriceSubTotal float64
	TaxSubTotal   float64
	GrandTotal    float64
	Products      []*Product
}
