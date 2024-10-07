package routes

import (
	"encoding/json"
	"fmt"
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
			r.Get("/", handler.GetThreads())
			r.Get("/{id}", handler.GetThread())
			r.Post("/", handler.CreateThread())
			r.Put("/{id}", handler.UpdateThread())
			r.Delete("/{id}", handler.DeleteThread())
		})
		r.Route("/threads/{id}/posts", func(r chi.Router) {
			r.Get("/", handler.GetPosts())
			r.Post("/", handler.CreatePost())
		})
		r.Route("/posts", func(r chi.Router) {
			r.Get("/{id}", handler.GetPost())
			r.Put("/{id}", handler.UpdatePost())
			r.Delete("/{id}", handler.DeletePost())
		})
	})

	return handler
}

type Handler struct {
	*chi.Mux
	store db.Store
}

/**
* * The GetThreads method returns a list of all threads.
* * The CreateThread method creates a new thread.
* * The GetThread method returns a single thread.
* * The UpdateThread method updates a thread.
* * The DeleteThread method deletes a thread.
 */
func (handler *Handler) GetThreads() http.HandlerFunc {
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

/**
* * The GetPosts method returns a list of all posts for a given thread.
* * The CreatePost method creates a new post for a given thread.
 */
func (handler *Handler) GetPosts() http.HandlerFunc {
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

		posts, err := handler.store.PostsByThread(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"posts": posts,
		}

		if posts == nil {
			response["posts"] = []types.Post{}
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (handler *Handler) CreatePost() http.HandlerFunc {
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

		var post types.Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		post.ThreadID = id
		if err := handler.store.CreatePost(&post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (handler *Handler) GetPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		if postID == "" {
			http.Error(w, "missing post ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(postID)
		if err != nil {
			http.Error(w, "invalid thread ID", http.StatusBadRequest)
			return
		}

		post, err := handler.store.Post(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"post": post,
		}

		fmt.Println(post)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (handler *Handler) UpdatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		if postID == "" {
			http.Error(w, "missing post ID", http.StatusBadRequest)
			return
		}

		var post types.Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(postID)
		if err != nil {
			http.Error(w, "invalid post ID", http.StatusBadRequest)
			return
		}
		post.ID = id
		if err := handler.store.UpdatePost(&post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (handler *Handler) DeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		if postID == "" {
			http.Error(w, "missing post ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(postID)
		if err != nil {
			http.Error(w, "invalid post ID", http.StatusBadRequest)
			return
		}

		if err := handler.store.DeletePost(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
