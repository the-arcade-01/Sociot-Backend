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
		router.Mount("/greet", greetRoutes())
		router.Mount("/users", userRoutes(server.AppConfig))
		router.Mount("/posts", postRoutes())
		router.Mount("/comments", commentRoutes())
	})
	server.Router.Get("/swagger/*", httpSwagger.WrapHandler)
	server.Router.Mount("/api", versionOne)
}

func greetRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", controller.Greet)
	return r
}

func userRoutes(appConfig *config.AppConfig) chi.Router {
	userRepo := repo.NewUserRepository(appConfig.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	claims := map[string]interface{}{"id": 1}
	_, token, _ := appConfig.Token.Encode(claims)
	log.Println(token)

	userRouter := chi.NewRouter()
	userRouter.Get("/", userController.GetUsers)
	userRouter.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(appConfig.Token))
		r.Use(jwtauth.Authenticator)
		r.Get("/{id}", userController.GetUserById)
	})

	return userRouter
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

// @title						Sociot Backend
// @version					1.0
// @description				REST API service written in Go.
//
// @contact.name				arcade
// @contact.url				https://github.com/the-arcade-01/Sociot-Backend
// @contact.email
//
// @host						localhost:5000
// @BasePath					/v1
func main() {
	appConfig := config.LoadConfig()
	server := CreateServer(appConfig)

	server.MountMiddlerwares()
	server.MountHandlers()

	log.Println("server running on port:5000")
	http.ListenAndServe(":5000", server.Router)
}
