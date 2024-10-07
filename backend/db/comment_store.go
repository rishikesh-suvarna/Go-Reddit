package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/rishikesh-suvarna/go-reddit/types"
)

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id int) (types.Comment, error) {
	var comment types.Comment
	err := s.Get(&comment, "SELECT * FROM comments WHERE id = $1", id)
	if err != nil {
		return types.Comment{}, err
	}
	return comment, nil
}

func (s *CommentStore) CommentsByPost(postID int) ([]types.Comment, error) {
	var comments []types.Comment
	err := s.Select(&comments, "SELECT * FROM comments WHERE post_id = $1", postID)
	if err != nil {
		return []types.Comment{}, err
	}
	return comments, nil
}

func (s *CommentStore) CreateComment(comment *types.Comment) error {
	err := s.Get(comment, "INSERT INTO comments (id, post_id, content, votes) VALUES ($1, $2, $3, $4) RETURNING *", comment.ID, comment.PostID, comment.Content, comment.Votes)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentStore) UpdateComment(comment *types.Comment) error {
	err := s.Get(comment, "UPDATE comments SET content = $1, votes = $2 WHERE id = $3 RETURNING *", comment.Content, comment.Votes, comment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentStore) DeleteComment(id int) error {
	_, err := s.Exec("DELETE FROM comments WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
