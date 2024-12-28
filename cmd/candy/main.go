package main

import (
	"candy/internal/app"
	"candy/pkg/loger"
	"context"
	"os"
	"os/signal"
)

func main() {
	apps := app.Init()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := apps.Run(ctx); err != nil {
		loger.New().Error().Err(err)
	}
}
