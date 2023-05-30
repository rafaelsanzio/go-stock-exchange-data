package main

import (
	"context"

	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/config/key"
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/scheduler"
)

func main() {
	ctx := context.Background()

	key.LoadEnvVars()

	scheduler.Run(ctx)
}
