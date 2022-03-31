package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

const PYTHON_PROGRAM_FILE string = "program.py"
type Program struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
    Code string `json:"code"`
	Flow string `json:"flow"`
}

func Route() *chi.Mux {
	mux := chi.NewMux()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300, 
	})
	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler,
		setResponseHeadersMiddleware,
	)

	mux.Route("/programs", func(r chi.Router) {
		r.Get("/", getAllPrograms)
		r.Post("/", saveProgram)
		r.Post("/execute", executeProgram)
		r.Route("/{uid}", func(r chi.Router) {
			r.Get("/", getProgramById)
		})
	})

	return mux
}

func setResponseHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func hasError(err error, w http.ResponseWriter, msg string, statusCode int) bool {
	hasErr := err != nil
	if hasErr {
		errMsg := fmt.Errorf("Error: %s", msg).Error()
		fmt.Print(errMsg)
		http.Error(w, errMsg, statusCode)
	}

	return hasErr
}