package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mohamedsamara/go-vue-auth/handlers"
)

func APIRoutes(h *handlers.BaseHandler) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("pong")
		})
	})

	r.Mount("/auth", AuthRoutes(h))
	r.Mount("/user", UserRoutes(h))

	return r
}
