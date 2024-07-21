package pg

import (
	"context"
	"database/sql"
	"example/store"
)

type ProductImplDB struct {
	db *sql.DB
}

func NewProductImplDB(db *sql.DB) *ProductImplDB {
	return &ProductImplDB{
		db: db,
	}
}

const insertProductSQL = `
	INSERT INTO products 
	(product_id, product_name, product_stock, product_price)
	VALUES ($1, $2, $3, $4);
`

func (p *ProductImplDB) Create(ctx context.Context, Product *store.ProductData) error {
	row := p.db.QueryRowContext(ctx, insertProductSQL,
		Product.ProductID,
		Product.ProductName,
		Product.ProductStock,
		Product.ProductPrice,
	)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}
