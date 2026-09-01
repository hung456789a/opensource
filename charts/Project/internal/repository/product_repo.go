package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"Project/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetByID(ctx context.Context, productIDs []string) ([]domain.Product, error) {
	if len(productIDs) == 0 {
		return []domain.Product{}, nil
	}

	placeholders := make([]string, len(productIDs))
	args := make([]interface{}, len(productIDs))

	for i, id := range productIDs {
		placeholders[i] = "?"
		args[i] = id
	}
	///combinate a full query

	query := fmt.Sprintf(
		"SELECT id, name, price, stock FROM products WHERE id IN (%s)",
		strings.Join(placeholders, ","),
	)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// map each row to domain.Product
	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, fmt.Errorf("lỗi scan dòng product: %w", err)
		}
		products = append(products, p)
	}

	//

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("lỗi duyệt kết quả: %w", err)
	}
	return products, nil
}
