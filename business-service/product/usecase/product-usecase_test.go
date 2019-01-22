package usecase_test

import (
	"context"
	"testing"

	"github.com/riyan10dec/tax-calc/business-service/product/mocks"
	"github.com/riyan10dec/tax-calc/business-service/product/usecase"
	"github.com/riyan10dec/tax-calc/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStore(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := models.Product{
		Name:    "Product 1",
		Price:   1000,
		TaxCode: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Store", mock.Anything, mockProduct).Once().Return(nil, &mockProduct)
		u := usecase.NewProductUsecase(mockProductRepo)

		err, _ := u.Store(context.TODO(), mockProduct)
		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
	})
}
func TestCalculateTax(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	products := make([]*models.Product, 0)
	products = append(products, &models.Product{
		ID:          1,
		Name:        "Lucky Stretch",
		Price:       1000,
		ProductType: "Tobacco",
		Refundable:  false,
		TaxCode:     2,
		TaxPrice:    30,
		Total:       1030,
	})
	products = append(products, &models.Product{
		ID:          2,
		Name:        "Big Mac",
		Price:       1000,
		ProductType: "Food & Beverage",
		Refundable:  true,
		TaxCode:     1,
		TaxPrice:    100,
		Total:       1100,
	})
	products = append(products, &models.Product{
		ID:          3,
		Name:        "Movie",
		Price:       150,
		ProductType: "Entertainment",
		Refundable:  false,
		TaxCode:     3,
		TaxPrice:    0.5,
		Total:       150.5,
	})
	products = append(products, &models.Product{
		ID:          4,
		Name:        "Movie",
		Price:       50,
		ProductType: "Entertainment",
		Refundable:  false,
		TaxCode:     3,
		TaxPrice:    0,
		Total:       50,
	})
	mockProductSummary := &models.ProductSummary{
		GrandTotal:    2330.5,
		PriceSubTotal: 2200,
		TaxSubTotal:   130.5,
		Products:      products,
	}
	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("GetProducts", mock.Anything, mock.AnythingOfType("[]int32")).Return(nil, products).Once()
		u := usecase.NewProductUsecase(mockProductRepo)

		err, productSummary := u.CalculateTax(context.TODO(), []int32{1, 2, 3})

		assert.NoError(t, err)
		assert.NotNil(t, products)

		assert.Equal(t, mockProductSummary, productSummary)
		mockProductRepo.AssertExpectations(t)
	})

}
