package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"movie-api/business/posts"
	"movie-api/business/posts/model"
	"net/http"
)

type postsRepository struct {
	client  *http.Client
	baseUrl string
}

// NewRepository - init user repository
func NewRepository(client *http.Client, url string) posts.Repository {
	return &postsRepository{
		client:  client,
		baseUrl: url,
	}
}

func (r *postsRepository) FindByID(id string) (model.Posts, error) {
	var result model.Posts
	url := r.baseUrl + "posts"
	if len(id) > 0 {
		url = url + "/" + id
	}

	resp, err := r.client.Get(url)
	if err != nil {
		return result, err
	}
	json.NewDecoder(resp.Body).Decode(&result)
	defer resp.Body.Close()

	return result, nil
}

func (r *postsRepository) Getpost() ([]model.Posts, error) {
	var result []model.Posts
	url := r.baseUrl + "posts"

	resp, err := r.client.Get(url)
	if err != nil {
		return result, err
	}
	json.NewDecoder(resp.Body).Decode(&result)
	defer resp.Body.Close()

	return result, nil
}

func (r *postsRepository) CreatePost(data model.Posts) (model.Posts, error) {
	var result model.Posts
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("executePost err  (Marshal) :", err)
		return result, err
	}

	resp, err := r.client.Post(r.baseUrl+"posts", "application/json", bytes.NewReader(b))
	if err != nil {
		fmt.Println("executePost err (client.Post) :", err)
		return result, err
	}

	json.NewDecoder(resp.Body).Decode(&result)
	defer resp.Body.Close()

	return result, nil
}

func (r *postsRepository) UpdatePost(id string, data model.Posts) (model.Posts, error) {
	var result model.Posts
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("executePost err  (Marshal) :", err)
		return result, err
	}
	req, err := http.NewRequest(http.MethodPut, r.baseUrl+"posts/"+id, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := r.client.Do(req)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&result)
	defer resp.Body.Close()

	return result, nil
}
func (r *postsRepository) DeletePost(id string) (model.Posts, error) {
	var result model.Posts

	req, err := http.NewRequest("DELETE", r.baseUrl+""+id, nil)
	if err != nil {
		fmt.Println(err)
		return result, nil
	}
	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := r.client.Do(req)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&result)
	defer resp.Body.Close()
	return result, nil
}
func (r *postsRepository) SearchPost(id string) (model.Posts, error) {
	var u model.Posts
	return u, nil
}
