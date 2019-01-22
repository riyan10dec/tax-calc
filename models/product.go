package models

type Product struct {
	ID          int32
	Name        string
	Price       float64
	TaxCode     int32
	Refundable  bool
	TaxPrice    float64
	ProductType string
	Total       float64
}
