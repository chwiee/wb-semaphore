package main

import (
	"log"
	"net/http"
	"time"

	"github.com/chwiee/wb-semaphore/internal/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	h := handlers.NewProjectHandler()

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// Rotas
	r.Get("/projects", h.ListProjects)
	r.Get("/project/{id}", h.GetProjectDetail)

	// Health
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })

	addr := ":8000"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
