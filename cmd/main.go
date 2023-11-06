package main

import (
	"log"
	"net/http"
	controller "sociot/internal/controller"
	repo "sociot/internal/repository"
	service "sociot/internal/service"

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

func (server *Server) MountHandlers() {
	versionOne := chi.NewRouter()

	versionOne.Route("/v1", func(router chi.Router) {
		router.Mount("/greet", greetRoutes())
		router.Mount("/users", userRoutes())
		router.Mount("/posts", postRoutes())
		router.Mount("/comments", commentRoutes())
	})

	server.Router.Mount("/api", versionOne)
}

func greetRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", controller.Greet)
	return r
}

func userRoutes() chi.Router {
	userRepo := repo.NewUserRepository(nil)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := chi.NewRouter()
	r.Get("/", userController.GetUsers)

	return r
}

func postRoutes() chi.Router {
	postRepo := repo.NewPostRepository(nil)
	postService := service.NewPostService(postRepo)
	postController := controller.NewPostController(postService)

	r := chi.NewRouter()
	r.Get("/", postController.GetPosts)

	return r
}

func commentRoutes() chi.Router {
	commentRepo := repo.NewCommentRepository(nil)
	commentService := service.NewCommentService(commentRepo)
	commentController := controller.NewCommentController(commentService)

	r := chi.NewRouter()
	r.Get("/", commentController.GetCommentById)

	return r
}

func main() {
	server := CreateServer()

	server.MountMiddlerwares()
	server.MountHandlers()

	log.Println("server running on port:5000")
	http.ListenAndServe(":5000", server.Router)
}
