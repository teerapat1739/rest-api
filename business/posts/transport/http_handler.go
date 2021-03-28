package transport

import (
	"movie-api/business/posts"
	"movie-api/business/posts/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// PostsHTTPHandler - posts handler for http request

type PostsHTTPHandler struct {
	postsSerivce posts.Service
}

// NewPostsHTTPHandler - init posts
func NewPostsHTTPHandler(e *echo.Echo, postsSerivce posts.Service) {
	handler := PostsHTTPHandler{postsSerivce}

	api := e.Group("/api")
	api.Use(middleware.Logger())

	// Added authentication
	api.GET("/posts", handler.GetPosts)
	api.GET("/posts/:id", handler.FindByID)
	api.POST("/posts", handler.CreatePosts)
	api.PUT("/posts/:id", handler.UpdatePosts)
	api.DELETE("/posts/:id", handler.DeletePost)
}

func (h *PostsHTTPHandler) FindByID(c echo.Context) error {
	id := c.Param("id")

	res, err := h.postsSerivce.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, res)
}

func (h *PostsHTTPHandler) GetPosts(c echo.Context) error {

	res, err := h.postsSerivce.Getpost()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, res)
}

func (h *PostsHTTPHandler) CreatePosts(c echo.Context) error {
	var m model.Posts

	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	res, err := h.postsSerivce.CreatePost(m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, res)
}

func (h *PostsHTTPHandler) UpdatePosts(c echo.Context) error {
	var m model.Posts
	id := c.Param("id")
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	res, err := h.postsSerivce.UpdatePost(id, m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, res)
}

func (h *PostsHTTPHandler) DeletePost(c echo.Context) error {
	id := c.Param("id")

	res, err := h.postsSerivce.DeletePost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, res)
}
