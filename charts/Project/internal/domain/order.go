package domain

import (
	"context"
)

type CreateOrderInput struct {
	UserID    string
	ProductID string
	Quantity  int
}

type Order struct {
	ID          string
	UserID      string
	TotalAmount float64
	Status      int
}

type OrderItem struct {
	OrderID   string
	ProductID string
	Quantity  string
	Price     string
}

type Product struct {
	ID    string
	Name  string
	Price float64
}

type ProductRepository interface {
	GetProductByID(ctx context.Context, productIDs []string) ([]Product, error)
}
