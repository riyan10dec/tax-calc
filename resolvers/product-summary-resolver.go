package resolvers

import (
	"context"

	"github.com/riyan10dec/tax-calc/business-service/product"

	"github.com/riyan10dec/tax-calc/business-service/product/usecase"
	"github.com/riyan10dec/tax-calc/models"
)

type ProductSummaryResolver struct {
	m *models.ProductSummary
}

func (psr *ProductSummaryResolver) GrandTotal() float64 {
	return psr.m.GrandTotal
}
func (psr *ProductSummaryResolver) TaxSubTotal() float64 {
	return psr.m.TaxSubTotal
}
func (psr *ProductSummaryResolver) PriceSubTotal() float64 {
	return psr.m.PriceSubTotal
}

func (psr *ProductSummaryResolver) Products(ctx context.Context) (*[]*ProductResolver, error) {
	var resolvers = make([]*ProductResolver, 0, len(psr.m.Products))
	for _, p := range psr.m.Products {
		resolvers = append(resolvers, &ProductResolver{
			m: p,
		})
	}

	return &resolvers, nil
}

func (r *Resolver) CalculateTax(ctx context.Context, args struct {
	ProductIDs []int32
}) (*ProductSummaryResolver, error) {
	productRepo := ctx.Value("ProductRepository").(product.ProductRepository)
	bu := usecase.NewProductUsecase(productRepo)
	err, productSummary := bu.CalculateTax(ctx, args.ProductIDs)
	if err != nil {
		return nil, err
	}
	return &ProductSummaryResolver{
		m: productSummary,
	}, nil
}
