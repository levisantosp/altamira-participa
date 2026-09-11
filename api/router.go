package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/levisantosp/altamira-participa/api/routes/auth"
	"github.com/levisantosp/altamira-participa/api/routes/issues"
	"github.com/levisantosp/altamira-participa/api/routes/users"
	"github.com/levisantosp/altamira-participa/api/utils"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RedirectSlashes)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: utils.Env.TrustedOrigins,
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"PATCH",
			"OPTIONS",
		},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, world!"))
	})

	api := humachi.New(r, huma.DefaultConfig("api docs", "0.0.0"))

	auth.Routes(api)
	users.Routes(api)
	issues.Routes(api)

	return r
}
