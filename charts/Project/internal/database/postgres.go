package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

func Dbconnect(ctx context.Context, host string, port int, username string, password string, dbname string) (*sql.DB, error) {
	logger := slog.Default().With(
		slog.String("driver", "postgres"),
		slog.String("host", host),
		slog.Int("port", port),
		slog.String("user", username),
		slog.String("dbname", dbname),
	)

	logger.InfoContext(ctx, "đang kết nối database")
	start := time.Now()

	// Kết nối PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)
	pgDB, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.ErrorContext(ctx, "cấu hình kết nối không hợp lệ", slog.Any("error", err))
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	// Kiểm tra kết nối thành công
	if err := pgDB.PingContext(ctx); err != nil {
		logger.ErrorContext(ctx, "kết nối database thất bại",
			slog.Any("error", err),
			slog.Duration("elapsed", time.Since(start)),
		)
		pgDB.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	logger.InfoContext(ctx, "kết nối database thành công",
		slog.Duration("elapsed", time.Since(start)),
	)
	return pgDB, nil
}
