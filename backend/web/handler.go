package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rishikesh-suvarna/go-reddit/db"
	"github.com/rishikesh-suvarna/go-reddit/types"
)

func NewHandler(store db.Store) *Handler {
	handler := &Handler{
		Mux:   chi.NewRouter(),
		store: store,
	}

	handler.Use(middleware.Logger)
	handler.Route("/threads", func(r chi.Router) {
		r.Get("/", handler.ThreadsList())
		r.Post("/", handler.CreateThread())
	})

	return handler
}

type Handler struct {
	*chi.Mux
	store db.Store
}

func (handler *Handler) ThreadsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		threads, err := handler.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"threads": threads,
		}
		if threads == nil {
			response["threads"] = []types.Thread{}
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (handler *Handler) CreateThread() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var thread types.Thread
		if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := handler.store.CreateThread(&thread); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(thread); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
