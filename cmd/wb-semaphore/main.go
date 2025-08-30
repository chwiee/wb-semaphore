package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	println("Starting project service...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/projects", projectHandler.GetProjects)
	http.ListenAndServe(":8000", r)
}
