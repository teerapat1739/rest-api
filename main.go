package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subosito/gotenv"

	postsrepo "movie-api/business/posts/repository"
	postssrc "movie-api/business/posts/service"
	postshdlr "movie-api/business/posts/transport"
)

func init() {
	gotenv.Load()
}

func main() {

	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}

	e := echo.New()
	e.Use(middleware.Recover())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	postsrepo := postsrepo.NewRepository(client, "https://jsonplaceholder.typicode.com/")

	postssrc := postssrc.NewUserService(postsrepo)

	postshdlr.NewPostsHTTPHandler(e, postssrc)

	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
