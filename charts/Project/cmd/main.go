package main

import (
	"Project/internal/domain"
	"Project/internal/repository"
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	productRepo := repository.NewProductRepository()
	ids := []string{"P01", "P02"}
	products, err := productRepo.GetProductID(ctx, ids)
	if err != nil {
		log.Fatalf("Error when retrieve products: %v", err)
	}
	fmt.Println("--- Results ---")
	for _, p := range products {
		fmt.Printf("ID: %s, Name: %s, Price: %f\n", p.ID, p.Name, p.Price)
	}
	order := domain.Order{
		ID:          "ORD-1001",
		UserID:      "USER-01",
		TotalAmount: 1525.0,
		Status:      1,
	}
	fmt.Printf("\nĐã tạo đơn hàng thử nghiệm: %+v\n", order)

}
