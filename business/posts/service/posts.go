package service

import (
	"movie-api/business/posts"
	"movie-api/business/posts/model"
)

type postsService struct {
	postsRepo posts.Repository
}

// userService - init user service
func NewUserService(postsRepo posts.Repository) posts.Service {
	return &postsService{
		postsRepo: postsRepo,
	}
}

func (s *postsService) FindByID(id string) (model.Posts, error) {
	return s.postsRepo.FindByID(id)
}
func (s *postsService) Getpost() ([]model.Posts, error) {
	return s.postsRepo.Getpost()
}

func (r *postsService) CreatePost(data model.Posts) (model.Posts, error) {
	return r.postsRepo.CreatePost(data)
}
func (r *postsService) UpdatePost(id string, data model.Posts) (model.Posts, error) {

	return r.postsRepo.UpdatePost(id, data)
}
func (r *postsService) DeletePost(id string) (model.Posts, error) {

	return r.postsRepo.DeletePost(id)
}
func (r *postsService) SearchPost(id string) (model.Posts, error) {
	var u model.Posts
	return u, nil
}
