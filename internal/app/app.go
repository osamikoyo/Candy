package app

import (
	"candy/internal/config"
	"candy/internal/data"
	"candy/pkg/loger"
	"context"
	"fmt"
	"net/http"
)

type App struct {
	DB     *data.Database
	Logger *loger.Logger
	server *http.Server
}

func Init() *App {
	cfg := config.Get()

	db := data.New()
	logger := loger.New()
	return &App{
		DB:     db,
		Logger: &logger,
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		},
	}
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		a.Logger.Info().Msg("Server stopped!! :3")
		err := a.server.Shutdown(ctx)
		if err != nil {
			return
		}
	}()

	a.Logger.Info().Msg("Starting the Server!! :3")

	if err := a.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
