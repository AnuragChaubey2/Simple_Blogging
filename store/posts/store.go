package posts

import (
	"database/sql"
	"errors"
	"time"

	"github.com/AnuragChaubey2/Simple_Blogging.git/models"
	"github.com/AnuragChaubey2/Simple_Blogging.git/store/utils"
)

type Posts struct {
    db *sql.DB
}

func New(db *sql.DB) Posts {
    return Posts{db: db}
}

func (s Posts) CreatePost(post models.Post) (int, error) {
    currentTime := time.Now()
    query := "INSERT INTO posts (title, content, author, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
    result, err := s.db.Exec(query, post.Title, post.Content, post.Author, currentTime, currentTime)
    if err != nil {
        return 0, err
    }
    postID, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(postID), nil
}

func (s Posts) GetPostByID(id int) (models.Post, error) {
    var post models.Post
    query := "SELECT id, title, content, author, CAST(created_at AS CHAR), CAST(updated_at AS CHAR) FROM posts WHERE id = ?"
    var createdTimeStr, updatedTimeStr string

    err := s.db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.Author, &createdTimeStr, &updatedTimeStr)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return models.Post{}, err
        }
        return models.Post{}, err
    }

    if createdTimeStr != "" {
        createdTime, err := utils.ParseTime(createdTimeStr)
        if err != nil {
            return models.Post{}, err
        }
        post.CreatedAt = createdTime
    }

    if updatedTimeStr != "" {
        updatedTime, err := utils.ParseTime(updatedTimeStr)
        if err != nil {
            return models.Post{}, err
        }
        post.UpdatedAt = updatedTime
    }

    return post, nil
}

func (s Posts) GetAllPosts() ([]models.Post, error) {
    var posts []models.Post
    query := "SELECT id, title, content, author, CAST(created_at AS CHAR), CAST(updated_at AS CHAR) FROM posts"
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var post models.Post
        var createdTimeStr, updatedTimeStr string

        if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &createdTimeStr, &updatedTimeStr); err != nil {
            return nil, err
        }

        if createdTimeStr != "" {
            createdTime, err := utils.ParseTime(createdTimeStr)
            if err != nil {
                return nil, err
            }
            post.CreatedAt = createdTime
        }

        if updatedTimeStr != "" {
            updatedTime, err := utils.ParseTime(updatedTimeStr)
            if err != nil {
                return nil, err
            }
            post.UpdatedAt = updatedTime
        }

        posts = append(posts, post)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return posts, nil
}

func (s Posts) UpdatePost(id int, post models.Post) error {
    currentTime := time.Now()
    query := "UPDATE posts SET title = ?, content = ?, author = ?, created_at = ?, updated_at = ? WHERE id = ?"
    
    if post.CreatedAt.IsZero() {
        post.CreatedAt = currentTime
    }

    _, err := s.db.Exec(query, post.Title, post.Content, post.Author, post.CreatedAt, currentTime, id)
    if err != nil {
        return err
    }
    return nil
}

func (s Posts) DeletePost(id int) error {
    query := "DELETE FROM posts WHERE id = ?"
    _, err := s.db.Exec(query, id)
    if err != nil {
        return err
    }
    return nil
}