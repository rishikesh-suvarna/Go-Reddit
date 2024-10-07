package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/rishikesh-suvarna/go-reddit/types"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id int) (types.Post, error) {
	var post types.Post
	err := s.Get(&post, "SELECT * FROM posts WHERE id = $1", id)
	if err != nil {
		return types.Post{}, err
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
	err := s.Get(post, "INSERT INTO posts (id, thread_id, title, content, votes) VALUES ($1, $2, $3, $4, $5) RETURNING *", post.ID, post.ThreadID, post.Title, post.Content, post.Votes)
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
