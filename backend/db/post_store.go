package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/rishikesh-suvarna/go-reddit/types"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id int) (*types.Post, error) {
	// We'll use this to scan our row
	post := &types.Post{
		Thread: &types.Thread{}, // Initialize the Thread pointer
	}

	err := s.Get(post, `
		SELECT 
			p.id,
			p.thread_id,
			p.title,
			p.content,
			p.votes,
			t.id as "thread.id",
			t.title as "thread.title",
			t.description as "thread.description"
		FROM posts p
		LEFT JOIN threads t ON t.id = p.thread_id
		WHERE p.id = $1`, id)

	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, ErrNotFound
		// }
		return nil, err
	}

	return post, nil
}

func (s *PostStore) PostsByThread(threadID int) ([]types.Post, error) {
	var posts []types.Post
	err := s.Select(&posts, "SELECT * FROM posts WHERE thread_id = $1", threadID)
	if err != nil {
		return []types.Post{}, err
	}
	return posts, nil
}

func (s *PostStore) CreatePost(post *types.Post) error {
	err := s.Get(post, "INSERT INTO posts (thread_id, title, content, votes) VALUES ($1, $2, $3, $4) RETURNING *", post.ThreadID, post.Title, post.Content, post.Votes)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStore) UpdatePost(post *types.Post) error {
	err := s.Get(post, "UPDATE posts SET title = $1, content = $2, votes = $3 WHERE id = $4 RETURNING *", post.Title, post.Content, post.Votes, post.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStore) DeletePost(id int) error {
	_, err := s.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
