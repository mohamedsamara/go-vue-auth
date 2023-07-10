package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mohamedsamara/golang-vue/frontend"
)

func IndexRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/", frontend.IndexRoute)
		r.Get("/assets/*", frontend.StaticRoute)
		r.Get("/favicon.ico", frontend.FaviconRoute)
	})

	return r
}
