package posts

import "movie-api/business/posts/model"

// Service - interface
type Service interface {
	Getpost() ([]model.Posts, error)
	FindByID(id string) (model.Posts, error)
	CreatePost(id model.Posts) (model.Posts, error)
	UpdatePost(id string, data model.Posts) (model.Posts, error)
	DeletePost(id string) (model.Posts, error)
	SearchPost(id string) (model.Posts, error)
}
