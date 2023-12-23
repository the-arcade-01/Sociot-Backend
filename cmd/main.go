package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"sociot/config"
	_ "sociot/docs"
	handler "sociot/internal/handler"
	repo "sociot/internal/repository"
	service "sociot/internal/service"
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
		votesRepo := repo.NewVotesRepo(server.AppConfig.DB)
		votesService := service.NewVotesService(votesRepo)
		votesHandler := handler.NewVotesHandler(votesService)

		votesRouter := chi.NewRouter()
		votesRouter.Get("/{postId}", votesHandler.GetVotesCountById)
		votesRouter.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(server.AppConfig.Token))
			r.Use(jwtauth.Authenticator)
			r.Put("/", votesHandler.UpdatePostVotesById)
			r.Get("/status", votesHandler.GetUserVoted)
		})

		postRepo := repo.NewPostRepository(server.AppConfig.DB, votesRepo)
		postService := service.NewPostService(postRepo, server.AppConfig.Token)
		postHandler := handler.NewPostHandler(postService)

		postRouter := chi.NewRouter()
		postRouter.Get("/", postHandler.GetPosts)
		postRouter.Get("/{id}", postHandler.GetPostById)
		postRouter.Get("/tags", postHandler.GetTags)
		postRouter.Put("/views/{id}", postHandler.UpdatePostViewsById)
		postRouter.Get("/users/{id}", postHandler.GetUserPosts)
		postRouter.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(server.AppConfig.Token))
			r.Use(jwtauth.Authenticator)
			r.Post("/", postHandler.CreatePost)
			r.Put("/{id}", postHandler.UpdatePostById)
			r.Delete("/{id}", postHandler.DeletePostById)
		})

		userRepo := repo.NewUserRepository(server.AppConfig.DB, postRepo)
		userService := service.NewUserService(userRepo, server.AppConfig.Token)
		userHandler := handler.NewUserHandler(userService)

		userRouter := chi.NewRouter()
		userRouter.Get("/", userHandler.GetUsers)
		userRouter.Post("/", userHandler.CreateUser)
		userRouter.Post("/login", userHandler.LoginUser)
		userRouter.Get("/stats/{id}", userHandler.GetUserStats)
		userRouter.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(server.AppConfig.Token))
			r.Use(jwtauth.Authenticator)
			r.Get("/{id}", userHandler.GetUserById)
			r.Put("/{id}", userHandler.UpdateUserById)
			r.Put("/password/{id}", userHandler.UpdateUserPasswordById)
			r.Delete("/{id}", userHandler.DeleteUserById)
		})

		generalRepo := repo.NewGeneralRepository(server.AppConfig.DB)
		generalService := service.NewGeneralService(generalRepo)
		generalHandler := handler.NewGeneralHandler(generalService)

		generalRouter := chi.NewRouter()
		generalRouter.Get("/", generalHandler.Search)

		router.Mount("/search", generalRouter)
		router.Mount("/users", userRouter)
		router.Mount("/posts", postRouter)
		router.Mount("/votes", votesRouter)
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
