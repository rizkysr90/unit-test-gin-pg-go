package store

import "context"

type ProductData struct {
	ProductID    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductStock int    `json:"product_stock"`
	ProductPrice int    `json:"product_price"`
}

type Product interface {
	Create(ctx context.Context, Product *ProductData) error
}
