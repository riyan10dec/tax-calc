package usecase

import (
	"context"
	"time"

	"github.com/riyan10dec/tax-calc/business-service/product"
	"github.com/riyan10dec/tax-calc/models"
)

type ProductUsecase struct {
	ProductRepository product.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUsecase(productRepo product.ProductRepository) product.ProductUsecase {
	return &ProductUsecase{
		ProductRepository: productRepo,
		contextTimeout:    time.Second * 60,
	}
}

func (p *ProductUsecase) CalculateTax(ctx context.Context, productIDs []int32) (error, *models.ProductSummary) {

	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	err, products := p.ProductRepository.GetProducts(ctx, productIDs)
	if err != nil {
		return err, nil
	}

	totalTax := 0.0
	totalPrice := 0.0
	grandTotal := 0.0
	for _, product := range products {
		if product.TaxCode == 1 {
			product.ProductType = "Food & Beverage"
			product.Refundable = true
			product.TaxPrice = 0.1 * product.Price
		}
		if product.TaxCode == 2 {
			product.ProductType = "Tobacco"
			product.Refundable = false
			product.TaxPrice = 10 + (0.02 * product.Price)
		}
		if product.TaxCode == 3 {
			product.ProductType = "Entertainment"
			product.Refundable = false
			if product.Price >= 100 {
				product.TaxPrice = 0.01 * (product.Price - 100)
			}
		}
		product.Total = product.Price + product.TaxPrice
		totalTax += product.TaxPrice
		totalPrice += product.Price
		grandTotal += product.Total
	}
	return nil, &models.ProductSummary{
		GrandTotal:    grandTotal,
		PriceSubTotal: totalPrice,
		TaxSubTotal:   totalTax,
		Products:      products,
	}
}

func (p *ProductUsecase) Store(ctx context.Context, product models.Product) (error, *models.Product) {

	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	err, productResult := p.ProductRepository.Store(ctx, product)
	if err != nil {
		return err, nil
	}
	return nil, productResult
}
