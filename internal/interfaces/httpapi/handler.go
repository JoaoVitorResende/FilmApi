package httpapi

import (
	"net/http"

	"main/internal/application"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Handler(uc application.SearchFilms) http.Handler {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", Get(uc))

	return r
}

func Get(uc application.SearchFilms) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		films, err := uc.Execute(r.Context(), r.URL.Query().Get("s"))
		if err != nil {
			handleError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, Response{Data: films})
	}
}