package main

import (
	"context"

	"github.com/FlezzProject/platform-api/internal/initializers"
)

func main() {
  ctx := context.Background()
  if err := initializers.InitializeGateway(ctx); err != nil {
    panic(err)
  }
}
