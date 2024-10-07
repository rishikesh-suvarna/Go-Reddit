# Go-Reddit

Go-Reddit is a platform similar to Reddit and Threads by Meta. It allows users to create threads, posts, and comments. The backend is built using Go and the Chi router.

## Features

- Create and manage threads
- Post and comment on threads
- User authentication and authorization

## Getting Started

### Prerequisites

- Go 1.20 or higher
- PostgreSQL

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/rishikesh-suvarna/Go-Reddit.git
    cd go-reddit
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run backend/cmd/internals/main.go
    ```

## API Endpoints

### Threads

- `GET /threads` - List all threads
- `POST /threads` - Create a new thread

### Posts

- `GET /threads/{id}/posts` - List all posts in a thread
- `POST /threads/{id}/posts` - Create a new post in a thread

### Comments

- `GET /posts/{id}/comments` - List all comments on a post
- `POST /posts/{id}/comments` - Create a new comment on a post

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
