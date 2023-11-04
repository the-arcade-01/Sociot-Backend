package main

import (
	"log"
	"net/http"
	Controller "sociot/internal/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	Router *chi.Mux
}

func CreateServer() *Server {
	server := &Server{}
	server.Router = chi.NewRouter()
	return server
}

func (server *Server) MountHandlers() {
	server.Router.Get("/greet", Controller.Greet)
}

func (server *Server) MountMiddlerwares() {
	server.Router.Use(middleware.Logger)
	server.Router.Use(middleware.CleanPath)
	server.Router.Use(middleware.Heartbeat("/"))
	server.Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

func main() {
	server := CreateServer()

	server.MountMiddlerwares()
	server.MountHandlers()

	log.Println("server running on port:5000")
	http.ListenAndServe(":5000", server.Router)
}
