package bootstrap

import (
	"context"
	"database/sql"
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
