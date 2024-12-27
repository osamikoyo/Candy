package handler

import (
	"candy/internal/data"
	"candy/pkg/loger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	db data.PostStorage
}

type handleFunc func(w http.ResponseWriter, r *http.Request) error

func handler(h handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			loger.New().Error().Err(err)
		}
	}
}

func (h Handler) RegisterHandlers(r *chi.Mux) {
	r.Post("/api/post/add", handler(h.AddPostHandler))
}
