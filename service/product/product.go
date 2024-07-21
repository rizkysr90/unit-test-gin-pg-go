package product

import (
	"context"
	"database/sql"
	"errors"
	"example/store"
	"strings"

	"github.com/google/uuid"
)

type ProductService struct {
	db           *sql.DB
	storeProduct store.Product
}

func NewProductService(db *sql.DB, storeProduct store.Product) *ProductService {
	return &ProductService{
		db:           db,
		storeProduct: storeProduct,
	}
}

type reqCreateProduct struct {
	*store.ProductData
}

func (req *reqCreateProduct) sanitize() {
	req.ProductName = strings.TrimSpace(req.ProductName)
}
func (req *reqCreateProduct) validate() error {
	if len(req.ProductName) == 0 {
		return errors.New("product_name is required")
	}
	return nil
}
func (p *ProductService) CreateProduct(ctx context.Context, data *store.ProductData) error {
	input := reqCreateProduct{data}

	input.sanitize()
	if err := input.validate(); err != nil {
		return err
	}
	insertedData := &store.ProductData{
		ProductID:    uuid.NewString(),
		ProductName:  data.ProductName,
		ProductStock: data.ProductStock,
		ProductPrice: data.ProductStock,
	}
	if err := p.storeProduct.Create(ctx, insertedData); err != nil {
		return err
	}
	return nil
}
