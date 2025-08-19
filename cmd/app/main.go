package main

import (
	"context"

	"github.com/solumD/go-service-template/internal/app"
)

func main() {
	ctx := context.Background()
	app.InitAndRun(ctx)
}
