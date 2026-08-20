package main

import (
	"context"
	"fmt"
)

func contextExample() {
	fmt.Println("Hello, World!")
	ctx := context.Background()
	fmt.Println(ctx)
}
