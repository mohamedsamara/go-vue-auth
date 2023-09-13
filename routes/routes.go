package routes

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mohamedsamara/go-vue-auth/constants"
	"github.com/mohamedsamara/go-vue-auth/handlers"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *chi.Mux {
	r := initChi()

	r.Use(middleware.Logger, middleware.WithValue(constants.DB_CONTEXT, db))

	h := handlers.New(db)

	r.Mount("/", IndexRoutes())
	r.Mount(constants.API_URL, APIRoutes(&h))

	return r
}

func initChi() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-User", "authorization", "x-jwt"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
	r.Use(middleware.Timeout(60 * time.Second))
	return r
}
