package types

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
	Thread(id int) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(thread *Thread) error
	UpdateThread(thread *Thread) error
	DeleteThread(id int) error
}

type PostStore interface {
	Post(id int) (Post, error)
	PostsByThread(threadID int) ([]Post, error)
	CreatePost(post *Post) error
	UpdatePost(post *Post) error
	DeletePost(id int) error
}

type CommentStore interface {
	Comment(id int) (Comment, error)
	CommentsByPost(postID int) ([]Comment, error)
	CreateComment(comment *Comment) error
	UpdateComment(comment *Comment) error
	DeleteComment(id int) error
}
