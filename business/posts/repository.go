package posts

import "movie-api/business/posts/model"

// Repository - interface
type Repository interface {
	Getpost() ([]model.Posts, error)
	FindByID(id string) (model.Posts, error)
	CreatePost(id model.Posts) (model.Posts, error)
	UpdatePost(id string, data model.Posts) (model.Posts, error)
	DeletePost(id string) (model.Posts, error)
	SearchPost(id string) (model.Posts, error)
}
