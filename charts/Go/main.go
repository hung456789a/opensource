package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Microsecond)
	defer cancel()

	go cookPho(ctx, status)
	go cookPizza(ctx, status)

	// do something
	for {
		select {
		case <-
		}
	}
}

func cookPho(ctx context.Context, status string) {
	select {
	case <-ctx.Done():
		fmt.Println("Cook pho done")
		return
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Cooking pho")
		return status
	}
}

func cookPizza(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Cook pizza done")
		return
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Cooking pizza")
		return status
	}
}
