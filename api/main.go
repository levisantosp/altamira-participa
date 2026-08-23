package main

import (
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/levisantosp/altamira-participa/api/db"
	"github.com/levisantosp/altamira-participa/api/routes/auth"
	"github.com/levisantosp/altamira-participa/api/routes/users"
	"github.com/levisantosp/altamira-participa/api/utils"

	_ "github.com/levisantosp/altamira-participa/api/ent/generated/runtime"
)

func main() {
	utils.LoadEnv()
	db.Connect()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
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
		w.Write([]byte("Hello, world!"))
	})

	api := humachi.New(r, huma.DefaultConfig("api docs", "0.0.0"))

	auth.Routes(api)
	users.Routes(api)

	log.Println("HTTP server running at http://localhost:3333")
	http.ListenAndServe(":3333", r)
}
