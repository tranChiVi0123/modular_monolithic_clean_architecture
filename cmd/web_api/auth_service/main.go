package main

import (
	"context"

	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/initializers"
)

func main() {
	ctx := context.Background()
	if err := initializers.InitializeAuthService(ctx); err != nil {
		panic(err)
	}
}
