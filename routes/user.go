package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mohamedsamara/go-vue-auth/auth"
	"github.com/mohamedsamara/go-vue-auth/handlers"
)

func UserRoutes(h *handlers.BaseHandler) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(auth.WithJWT)
		r.Get("/me", h.GetUser)
	})

	return r
}
