package resolvers

import (
	"context"

	"github.com/riyan10dec/tax-calc/business-service/product"
	"github.com/riyan10dec/tax-calc/business-service/product/usecase"
	"github.com/riyan10dec/tax-calc/models"
)

type ProductResolver struct {
	m *models.Product
}

func (pr *ProductResolver) ID() int32 {
	return pr.m.ID
}
func (pr *ProductResolver) Name() string {
	return pr.m.Name
}
func (pr *ProductResolver) Price() float64 {
	return pr.m.Price
}
func (pr *ProductResolver) TaxCode() int32 {
	return pr.m.TaxCode
}
func (pr *ProductResolver) Refundable() bool {
	return pr.m.Refundable
}
func (pr *ProductResolver) TaxPrice() float64 {
	return pr.m.TaxPrice
}
func (pr *ProductResolver) ProductType() string {
	return pr.m.ProductType
}
func (pr *ProductResolver) Total() float64 {
	return pr.m.Total
}

// ProductInput : Model for Insert new Product
type ProductInput struct {
	Name    string
	Price   float64
	TaxCode int32
}

func (r *Resolver) InputProduct(ctx context.Context, args *struct {
	Product ProductInput
}) (*ProductResolver, error) {
	productModel := models.Product{
		Name:    args.Product.Name,
		Price:   args.Product.Price,
		TaxCode: args.Product.TaxCode,
	}
	productRepo := ctx.Value("ProductRepository").(product.ProductRepository)
	bu := usecase.NewProductUsecase(productRepo)
	err, productResult := bu.Store(ctx, productModel)
	if err != nil {
		return nil, err
	}
	return &ProductResolver{productResult}, nil
}
