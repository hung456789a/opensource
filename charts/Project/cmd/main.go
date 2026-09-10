package main

import (
	"context"
	"log/slog"
	"os"

	"Project/internal/bootstrap"
	"Project/internal/database"
	"Project/internal/logging"
)

// func main() {
// 	ctx := context.Background()

// 	connStr := "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()
// 	productRepo := repository.NewProductRepository(db)
// 	ids := []string{"P01", "P02"}
// 	products, err := productRepo.GetByID(ctx, ids)
// 	if err != nil {
// 		log.Fatalf("Error when retrieve products: %v", err)
// 	}
// 	fmt.Println("--- Results ---")
// 	for _, p := range products {
// 		fmt.Printf("ID: %s, Name: %s, Price: %f\n", p.ID, p.Name, p.Price)
// 	}
// 	order := domain.Order{
// 		ID:          "ORD-1001",
// 		UserID:      "USER-01",
// 		TotalAmount: 1525.0,
// 		Status:      1,
// 	}s
// 	fmt.Printf("\nĐã tạo đơn hàng thử nghiệm: %+v\n", order)

// }

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slog.SetDefault(slog.New(logging.ContextHandler{Handler: jsonHandler}))

	ctx := logging.WithTraceID(context.Background(), logging.NewTraceID())

	db, err := database.Dbconnect(ctx, "localhost", 5432, "myuser", "mypassword", "postgres")
	if err != nil {
		slog.ErrorContext(ctx, "không thể khởi tạo database", slog.Any("error", err))
		os.Exit(1)
	}
	defer db.Close()

	exists, err := bootstrap.TableExists(db, ctx, "public", "users")
	if err != nil {
		slog.ErrorContext(ctx, "không thể kiểm tra bảng", slog.Any("error", err))
		os.Exit(1)
	}
	slog.InfoContext(ctx, "Kiểm tra bảng", slog.Bool("exists", exists))
	
	if !exists {
		err = bootstrap.CreateTable(db, ctx, "public", "users")
		if err != nil {
			slog.ErrorContext(ctx, "không thể tạo bảng", slog.Any("error", err))
			os.Exit(1)
		}
		slog.InfoContext(ctx, "Tạo bảng thành công")
	}
}
