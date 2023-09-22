package main

import (
	"context"

	"github.com/AnuragChaubey2/Simple_Blogging.git/models"
)

type PostsService interface {
	GetPostByID(ctx context.Context, id int) (models.Post, error)
	PostCreatePost(ctx context.Context, post models.Post) (int, error)
	UpdatePost(ctx context.Context, id int, post models.Post) error
	DeletePost(ctx context.Context, id int) error
}
