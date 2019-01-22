package repository

import (
	"context"
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/riyan10dec/tax-calc/business-service/product"
	"github.com/riyan10dec/tax-calc/models"
)

type sqlProductRepository struct {
	Conn *sql.DB
}

func NewSqlProductRepository(Conn *sql.DB) product.ProductRepository {
	return &sqlProductRepository{Conn}
}

func (s *sqlProductRepository) GetProducts(ctx context.Context, productIDs []int32) (error, []*models.Product) {
	query := `
	SELECT  
		p.ID,
		p.Name,
		p.Price,
		t.Code
	FROM Product p
		INNER JOIN Tax t on p.TaxID = t.ID
	WHERE p.ID IN ([params])
		`

	paramValues := make([]interface{}, 0)
	params := make([]string, 0)
	for _, productID := range productIDs {
		params = append(params, "?")
		paramValues = append(paramValues, productID)
	}
	query = strings.Replace(
		query,
		"[params]",
		strings.Join(params, ","),
		1,
	)
	rows, err := s.Conn.Query(query, paramValues...)
	if err != nil {
		return err, nil
	}

	defer rows.Close()

	result := make([]*models.Product, 0)
	for rows.Next() {
		p := new(models.Product)
		err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.TaxCode,
		)
		if err != nil {
			return err, nil
		}
		result = append(result, p)
	}
	if err != nil {
		return err, nil
	}
	return nil, result
}

func (s *sqlProductRepository) Store(ctx context.Context, p models.Product) (error, *models.Product) {
	insertProductQuery := `INSERT Product Set
	Name = ?,
	Price = ?,
	TaxID = 
		(SELECT ID 
		FROM Tax
		WHERE Code = ?)`
	_, err := s.Conn.ExecContext(ctx, insertProductQuery,
		p.Name,
		p.Price,
		p.TaxCode,
	)
	if err != nil {
		return err, nil
	}
	return nil, &p
}
