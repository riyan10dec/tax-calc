package product

import (
	"context"

	"github.com/riyan10dec/tax-calc/models"
)

type ProductUsecase interface {
	CalculateTax(ctx context.Context, productIDs []int32) (error, *models.ProductSummary)
	Store(ctx context.Context, product models.Product) (error, *models.Product)
}
