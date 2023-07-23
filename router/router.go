package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matheusjost/go-bettings-api/configs"
	"github.com/matheusjost/go-bettings-api/handlers"
)

func InitializeRouter() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	// TODO: add proper return of http codes and their status
	// TODO: add auth route and protect the others with middlewares
	// TODO: create a group base router specifying api version
	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetHTTPPort()), r)
}
