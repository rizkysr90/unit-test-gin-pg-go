package store

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockProduct is a mock implementation of the Product interface
type MockProduct struct {
	mock.Mock
}

// Create is the mock implementation of the Create method
func (m *MockProduct) Create(ctx context.Context, product *ProductData) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}
