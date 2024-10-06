package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pborman/uuid"
	"github.com/rishikesh-suvarna/go-reddit/types"
)

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id uuid.UUID) (types.Thread, error) {
	var thread types.Thread
	err := s.Get(&thread, "SELECT * FROM threads WHERE id = $1", id)
	if err != nil {
		return types.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}
	return thread, nil
}

func (s *ThreadStore) Threads() ([]types.Thread, error) {
	var threads []types.Thread
	err := s.Select(&threads, "SELECT * FROM threads")
	if err != nil {
		return []types.Thread{}, fmt.Errorf("error getting threads: %w", err)
	}
	return threads, nil
}

func (s *ThreadStore) CreateThread(thread *types.Thread) error {
	err := s.Get(thread, "INSERT INTO threads (title, description) VALUES ($1, $2) RETURNING *", thread.Title, thread.Description)
	if err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) UpdateThread(thread *types.Thread) error {
	err := s.Get(thread, "UPDATE threads SET title = $1, description = $2 WHERE id = $3 RETURNING *", thread.Title, thread.Description, thread.ID)
	if err != nil {
		return fmt.Errorf("error updating thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	_, err := s.Exec("DELETE FROM threads WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting thread: %w", err)
	}
	return nil
}
