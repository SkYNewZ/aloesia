package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SkYNewZ/aloesia/models"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAllPatientsHandler return all patient
func GetAllPatientsHandler(w http.ResponseWriter, r *http.Request) {
	u, err := models.GetAllPatients()
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(err) //nolint
		return
	}
	json.NewEncoder(w).Encode(&u) //nolint
}

// GetOnePatientHandler return one patient by given id
func GetOnePatientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	u, err := models.GetOnePatient(vars["id"])

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

// CreatePatientHandler creates patient
func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	json.NewDecoder(r.Body).Decode(&patient) //nolint

	// if given patient is invalid
	_, err := govalidator.ValidateStruct(&patient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewBadRequestError(http.StatusBadRequest, err.Error())) //nolint
		return
	}

	err = models.CreatePatient(&patient)

	// unexpected error
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(err) //nolint
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&patient) //nolint
}

// DeletePatientHandler delete patient by given id
func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := models.DeletePatient(vars["id"])

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
