package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	handler.Route("/api/v1", func(r chi.Router) {
		r.Route("/threads", func(r chi.Router) {
			r.Get("/", handler.GetThreadsList())
			r.Get("/{id}", handler.GetThread())
			r.Post("/", handler.CreateThread())
			r.Put("/{id}", handler.UpdateThread())
			r.Delete("/{id}", handler.DeleteThread())
		})
	})

	return handler
}

type Handler struct {
	*chi.Mux
	store db.Store
}

func (handler *Handler) GetThreadsList() http.HandlerFunc {
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

func (handler *Handler) GetThread() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		threadID := chi.URLParam(r, "id")
		if threadID == "" {
			http.Error(w, "missing thread ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(threadID)
		if err != nil {
			http.Error(w, "invalid thread ID", http.StatusBadRequest)
			return
		}

		thread, err := handler.store.Thread(id)
		if err != nil {
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

func (handler *Handler) UpdateThread() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		threadID := chi.URLParam(r, "id")
		if threadID == "" {
			http.Error(w, "missing thread ID", http.StatusBadRequest)
			return
		}

		var thread types.Thread
		if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(threadID)
		if err != nil {
			http.Error(w, "invalid thread ID", http.StatusBadRequest)
			return
		}
		thread.ID = id
		if err := handler.store.UpdateThread(&thread); err != nil {
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

func (handler *Handler) DeleteThread() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		threadID := chi.URLParam(r, "id")
		if threadID == "" {
			http.Error(w, "missing thread ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(threadID)
		if err != nil {
			http.Error(w, "invalid thread ID", http.StatusBadRequest)
			return
		}

		if err := handler.store.DeleteThread(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
