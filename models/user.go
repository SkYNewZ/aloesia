package models

import uuid "github.com/satori/go.uuid"

// User model struct
type User struct {
	ID                 uuid.UUID `json:"id"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Enable             bool      `json:"enable"`
	LastConnectionDate string    `json:"last_connection_date"`
}

// Users many users
type Users []User
