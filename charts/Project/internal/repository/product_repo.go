package repository

import (
	"context"
	"database/sql"
	"strings"
)

type ProductRepository interface {
	GetByIDs(ctx, productID []string)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetByIDs(ctx context.Context, productIDs []string) ([]domain.Product, error) {
	if len(productIDs) == 0 {
		return nil, nil
	}

	return nil, nil

	placeholders := make([]string, len(productIDs))
	args := make([]interface{}, len(productIDs))

	for i, id := range productIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	query := "SELECT * FROM products WHERE id IN (" + strings.Join(placeholders, ",") + ")"

	return nil, nil
}
