package product

import (
	"context"
	"example/store"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductSuccess(t *testing.T) {
	t.Run("it will success", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		mockProduct := store.MockProduct{}
		productService := NewProductService(db, &mockProduct)

		inputData := store.ProductData{
			ProductName:  "Indomie ",
			ProductStock: 12,
			ProductPrice: 3000,
		}
		testContext := context.Background()
		mockProduct.On("Create", testContext, mock.Anything).Return(nil)

		err = productService.CreateProduct(testContext, &inputData)
		assert.Nil(t, err, "it should be nil because it's success to insert new data")
	})
}
func TestCreateProductFailed(t *testing.T) {
	t.Run("it will failed because product name is empty", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		mockProduct := store.MockProduct{}
		productService := NewProductService(db, &mockProduct)

		inputData := store.ProductData{
			ProductName:  "",
			ProductStock: 12,
			ProductPrice: 3000,
		}
		testContext := context.Background()
		mockProduct.On("Create", testContext, mock.Anything).Return(nil)

		err = productService.CreateProduct(testContext, &inputData)
		assert.NotNil(t, err, "it should be not nil because there is an error")
	})
}
