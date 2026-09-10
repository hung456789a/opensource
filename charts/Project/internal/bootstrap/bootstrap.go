package bootstrap

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

func TableExists(db *sql.DB, ctx context.Context, schema, table string) (bool, error) {
	var exists bool
	
	query := `
		SELECT EXISTS (
			SELECT 1 FROM information_schema.tables
			WHERE table_schema = $1 AND table_name = $2
		)
	`
	
	err := db.QueryRowContext(ctx, query, schema, table).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func CreateTable(db *sql.DB, ctx context.Context, schema, table string) error {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s.%s (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`, pq.QuoteIdentifier(schema), pq.QuoteIdentifier(table))
	
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}