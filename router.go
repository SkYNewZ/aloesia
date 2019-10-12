package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SkYNewZ/aloesia/controllers"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s - %s", r.Method, r.RequestURI, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}

func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// define JSON as default return content type
		w.Header().Set("Content-type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

// InitializeRouter initialize main router app
func InitializeRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	apiRouter := r.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/health", controllers.HealthCheckHandler).Methods(http.MethodGet)

	if os.Getenv("CI") == "" {
		r.Use(loggingMiddleware)
	}
	r.Use(contentTypeMiddleware)
	return r
}
