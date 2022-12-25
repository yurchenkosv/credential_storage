package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/yurchenkosv/credential_storage/internal/controllers"
	"github.com/yurchenkosv/credential_storage/internal/repository"
	"github.com/yurchenkosv/credential_storage/internal/service"
)

func NewRouter(repo repository.Repository, tokenAuth *jwtauth.JWTAuth) chi.Router {
	var (
		authService = service.NewAuthService(repo)

		authController = controllers.NewAuthController(&authService, tokenAuth)
	)

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.StripSlashes)

	router.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.AllowContentType("application/json"))
			r.Post("/user/register", authController.HandleUserRegistration)
			r.Post("/user/login", authController.HanldeUserLogin)
		})
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator)
		})
	})
	return router
}
