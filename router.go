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
	apiRouter.HandleFunc("/login", controllers.LoginHandler).Methods(http.MethodPost)

	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", controllers.CreateUserHandler).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id}", controllers.GetUserHandler).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", controllers.DeleteUserHandler).Methods(http.MethodDelete)
	userRouter.HandleFunc("", controllers.GetAllUsersHandler).Methods(http.MethodGet)

	patientRouter := apiRouter.PathPrefix("/patients").Subrouter()
	patientRouter.HandleFunc("", controllers.GetAllPatientsHandler).Methods(http.MethodGet)
	patientRouter.HandleFunc("", controllers.CreatePatientHandler).Methods(http.MethodPost)
	patientRouter.HandleFunc("/{id}", controllers.GetOnePatientHandler).Methods(http.MethodGet)
	patientRouter.HandleFunc("/{id}", controllers.DeletePatientHandler).Methods(http.MethodDelete)

	// Disable http access log on testing
	if os.Getenv("CI") == "" {
		r.Use(loggingMiddleware)
	}
	r.Use(contentTypeMiddleware)
	return r
}
