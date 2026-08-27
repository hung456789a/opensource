package repository

import (
	"context"
	"database/sql"
	"domain"
	"fmt"
	"strings"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db: db}
}

// GetByIDs lấy danh sách sản phẩm theo danh sách productIDs truyền vào
func (r *productRepository) GetByIDs(ctx context.Context, productIDs []string) ([]domain.Product, error) {
	// 1. Kiểm tra nếu danh sách ID rỗng thì trả về ngay, không gọi DB
	if len(productIDs) == 0 {
		return []domain.Product{}, nil
	}

	// 2. Tạo chuỗi placeholder (?, ?, ?) tương ứng với số lượng productIDs
	// Ví dụ: len = 3 -> placeholders = "?,?,?"
	placeholders := make([]string, len(productIDs))
	args := make([]interface{}, len(productIDs))

	for i, id := range productIDs {
		placeholders[i] = "?" // Nếu dùng PostgreSQL thì đổi thành fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	// 3. Ghép thành câu lệnh SQL hoàn chỉnh
	query := fmt.Sprintf(
		"SELECT id, name, price, stock FROM products WHERE id IN (%s)",
		strings.Join(placeholders, ","),
	)

	// 4. Thực thi Query
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("lỗi query products: %w", err)
	}
	defer rows.Close()

	// 5. Duyệt từng dòng kết quả và Scan vào Struct
	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, fmt.Errorf("lỗi scan dòng product: %w", err)
		}
		products = append(products, p)
	}

	// 6. Kiểm tra lỗi sau khi kết thúc vòng lặp rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("lỗi duyệt rows product: %w", err)
	}

	return products, nil
}
