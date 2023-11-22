package main

import (
	"log"
	"net/http"
	"sociot/config"
	controller "sociot/internal/controller"
	repo "sociot/internal/repository"
	service "sociot/internal/service"

	_ "sociot/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
	AppConfig *config.AppConfig
	Router    *chi.Mux
}

func CreateServer(appConfig *config.AppConfig) *Server {
	server := &Server{}
	server.Router = chi.NewRouter()
	server.AppConfig = appConfig
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
		greetRouter := chi.NewRouter()
		greetRouter.Get("/greet", controller.Greet)

		postRepo := repo.NewPostRepository(server.AppConfig.DB)
		postService := service.NewPostService(postRepo, server.AppConfig.Token)
		postController := controller.NewPostController(postService)

		postRouter := chi.NewRouter()
		postRouter.Get("/", postController.GetPosts)
		postRouter.Get("/{id}", postController.GetPostById)
		postRouter.Put("/views/{id}", postController.UpdatePostViewsById)
		postRouter.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(server.AppConfig.Token))
			r.Use(jwtauth.Authenticator)
			r.Get("/users/{id}", postController.GetUserPosts)
			r.Post("/", postController.CreatePost)
			r.Put("/{id}", postController.UpdatePostById)
			r.Delete("/{id}", postController.DeletePostById)
		})

		userRepo := repo.NewUserRepository(server.AppConfig.DB, postRepo)
		userService := service.NewUserService(userRepo, server.AppConfig.Token)
		userController := controller.NewUserController(userService)

		userRouter := chi.NewRouter()
		userRouter.Get("/", userController.GetUsers)
		userRouter.Post("/", userController.CreateUser)
		userRouter.Post("/login", userController.LoginUser)
		userRouter.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(server.AppConfig.Token))
			r.Use(jwtauth.Authenticator)
			r.Get("/{id}", userController.GetUserById)
			r.Put("/{id}", userController.UpdateUserById)
			r.Put("/password/{id}", userController.UpdateUserPasswordById)
			r.Delete("/{id}", userController.DeleteUserById)
		})

		router.Mount("/greet", greetRouter)
		router.Mount("/users", userRouter)
		router.Mount("/posts", postRouter)
	})
	server.Router.Get("/swagger/*", httpSwagger.WrapHandler)
	server.Router.Mount("/api", versionOne)
}

// @title						Sociot Backend
// @version					1.0
// @description				REST API service written in Go.
//
// @contact.name				arcade
// @contact.url				https://github.com/the-arcade-01/Sociot-Backend
// @contact.email
//
// @host						localhost:5000
// @BasePath					/api/v1
func main() {
	appConfig := config.LoadConfig()
	server := CreateServer(appConfig)

	server.MountMiddlerwares()
	server.MountHandlers()

	log.Println("server running on port:5000")
	http.ListenAndServe(":5000", server.Router)
}
