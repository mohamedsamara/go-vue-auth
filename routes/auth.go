package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mohamedsamara/golang-vue/handlers"
)

func AuthRoutes(h *handlers.BaseHandler) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Post("/refresh", h.RefreshToken)
	})

	return r
}
