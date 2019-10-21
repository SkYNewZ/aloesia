package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/SkYNewZ/aloesia/models"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

// LoginHandler handle login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	// if given credentials are invalid
	_, err := govalidator.ValidateStruct(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewBadRequestError(http.StatusBadRequest, err.Error()))
		return
	}

	// find existing user
	userSearch, _ := models.FindUserByEmailAndPassword(creds.Email, creds.Password)

	if userSearch.Email == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// generate Jwt
	expirationTime := time.Now().Add(2 * time.Hour)

	var u models.JwtUser
	err = mapstructure.Decode(userSearch, &u)

	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(err)
		return
	}

	claims := &models.Claims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "api-dot-aloesia.appspot.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(models.JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
