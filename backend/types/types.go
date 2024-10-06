package types

import "github.com/pborman/uuid"

type Thread struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type Post struct {
	ID       int    `db:"id"`
	ThreadID int    `db:"thread_id"`
	Title    string `db:"title"`
	Content  string `db:"content"`
	Votes    int    `db:"votes"`
}

type Comment struct {
	ID      int    `db:"id"`
	PostID  int    `db:"post_id"`
	Content string `db:"content"`
	Votes   int    `db:"votes"`
}

type ThreadStore interface {
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(thread *Thread) error
	UpdateThread(thread *Thread) error
	DeleteThread(id uuid.UUID) error
}

type PostStore interface {
	Post(id uuid.UUID) (Post, error)
	PostsByThread(threadID uuid.UUID) ([]Post, error)
	CreatePost(post *Post) error
	UpdatePost(post *Post) error
	DeletePost(id uuid.UUID) error
}

type CommentStore interface {
	Comment(id uuid.UUID) (Comment, error)
	CommentsByPost(postID uuid.UUID) ([]Comment, error)
	CreateComment(comment *Comment) error
	UpdateComment(comment *Comment) error
	DeleteComment(id uuid.UUID) error
}
