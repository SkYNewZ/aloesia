package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/SkYNewZ/aloesia/models"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUserHandler create user from request JSON body
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user) //nolint

	// if given user is invalid
	_, err := govalidator.ValidateStruct(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewBadRequestError(http.StatusBadRequest, err.Error())) //nolint
		return
	}

	err = models.CreateUser(&user)
	// if user already exist
	if err != nil && err.Error() == "User already exist" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewBadRequestError(http.StatusBadRequest, err.Error())) //nolint
		return
	}

	// unexpected error
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(err) //nolint
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user) //nolint
}

// GetAllUsersHandler return all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	u, err := models.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(err) //nolint
		return
	}
	json.NewEncoder(w).Encode(&u) //nolint
}

// GetUserHandler return user by id
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	u, err := models.GetUser(vars["id"])

	// if no result found with this id
	if status.Code(err) == codes.NotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(err) //nolint
		return
	}

	json.NewEncoder(w).Encode(&u) //nolint
}

// DeleteUserHandler return user by id
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := models.DeleteUser(vars["id"])

	// if no result found with this id
	if status.Code(err) == codes.NotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(err) //nolint
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
