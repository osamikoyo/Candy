package app

import (
	"candy/internal/config"
	"candy/internal/data"
	"candy/internal/handler"
	"candy/pkg/loger"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type App struct {
	DB      *data.Database
	Logger  *loger.Logger
	server  *http.Server
	handler *handler.Handler
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
		handler: &handler.Handler{DB: db},
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

	r := chi.NewRouter()
	a.server.Handler = r

	a.Logger.Info().Msg("Starting the Server!! :3")
	a.handler.RegisterHandlers(r)

	if err := a.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
