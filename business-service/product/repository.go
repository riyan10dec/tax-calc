package product

import (
	"context"

	"github.com/riyan10dec/tax-calc/models"
)

type ProductRepository interface {
	GetProducts(ctx context.Context, productIDs []int32) (error, []*models.Product)
	Store(ctx context.Context, product models.Product) (error, *models.Product)
}
