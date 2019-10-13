package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var router *mux.Router
var rr *httptest.ResponseRecorder

func init() {
	router = InitializeRouter()
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr = httptest.NewRecorder()
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json; charset=utf-8"
	if rr.Header().Get("Content-type") != expectedContentType {
		t.Errorf("handler returned wrong content-type: got %v want %v", rr.Header().Get("Content-type"), expectedContentType)
	}
}
