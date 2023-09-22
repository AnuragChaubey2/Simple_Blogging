package posts

import (
    "context"

    "github.com/AnuragChaubey2/Simple_Blogging.git/models"
    "github.com/AnuragChaubey2/Simple_Blogging.git/store/posts"
)

type PostsService interface {
    GetPostByID(ctx context.Context, id int) (models.Post, error)
    GetAllPosts(ctx context.Context) ([]models.Post, error)
    CreatePost(ctx context.Context, post models.Post) (int, error)
    UpdatePost(ctx context.Context, id int, post models.Post) error
    DeletePost(ctx context.Context, id int) error
}

type Posts struct {
    postStore posts.Posts 
}

func NewPostsService(store posts.Posts) PostsService {
    return &Posts{postStore: store}
}

func (s *Posts) GetPostByID(ctx context.Context, id int) (models.Post, error) {
    post, err := s.postStore.GetPostByID(id)
    if err != nil {
        return models.Post{}, err
    }
    return post, nil
}


func (s *Posts) GetAllPosts(ctx context.Context) ([]models.Post, error) {

    return s.postStore.GetAllPosts()
}

func (s *Posts) CreatePost(ctx context.Context, post models.Post) (int, error) {

    if err := post.ValidatePost(); err != nil {
        return 0, err
    }

    return s.postStore.CreatePost(post)
}

func (s *Posts) UpdatePost(ctx context.Context, id int, post models.Post) error {

    if err := post.ValidatePost(); err != nil {
        return err
    }
    
    return s.postStore.UpdatePost(id, post)
}

func (s *Posts) DeletePost(ctx context.Context, id int) error {
    return s.postStore.DeletePost(id)
}
