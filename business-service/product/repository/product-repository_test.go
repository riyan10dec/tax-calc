package repository_test

import (
	"context"
	"testing"

	"github.com/riyan10dec/tax-calc/business-service/product/repository"
	"github.com/riyan10dec/tax-calc/models"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestStore(t *testing.T) {
	ar := &models.Product{
		Name:    "Product 1",
		Price:   1000.00,
		TaxCode: 1,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := `[INSERT Product Set
		Name = \\?,
		Price = \\?,
		TaxID = 
			(SELECT ID 
			FROM Tax
			WHERE Code = \\?)]`
	// prep := mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(ar.Name, ar.Price, ar.TaxCode).WillReturnResult(sqlmock.NewResult(0, 1))

	a := repository.NewSqlProductRepository(db)

	err, _ = a.Store(context.TODO(), *ar)
	assert.NoError(t, err)
	assert.Equal(t, int32(0), ar.ID)
}

func TestGetProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockProducts := []models.Product{
		models.Product{
			ID:      1,
			Name:    "Product 1",
			Price:   1500,
			TaxCode: 1,
		},
		models.Product{
			ID:      2,
			Name:    "Product 2",
			Price:   3000,
			TaxCode: 2,
		},
		models.Product{
			ID:      3,
			Name:    "Product 3",
			Price:   3000,
			TaxCode: 3,
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Price", "TaxID"}).
		AddRow(mockProducts[0].ID, mockProducts[0].Name, mockProducts[0].Price, mockProducts[0].TaxCode).
		AddRow(mockProducts[1].ID, mockProducts[1].Name, mockProducts[1].Price, mockProducts[1].TaxCode).
		AddRow(mockProducts[2].ID, mockProducts[2].Name, mockProducts[2].Price, mockProducts[2].TaxCode)

	query := `
	[SELECT  
		p.ID,
		p.Name,
		p.Price,
		t.Code
	FROM Product p
		INNER JOIN Tax t on p.TaxID = t.ID
	WHERE p.ID IN (1,2,3)]`

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := repository.NewSqlProductRepository(db)
	err, list := a.GetProducts(context.TODO(), []int32{1, 2, 3})
	assert.NoError(t, err)
	assert.Len(t, list, 3)
}
