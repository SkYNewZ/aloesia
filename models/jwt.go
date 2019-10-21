package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims content of JWT's payload
type Claims struct {
	User JwtUser `json:"user"`
	jwt.StandardClaims
}

// JwtUser user in JWT
type JwtUser struct {
	ID    string `json:"id" mapstructure:"id"`
	Email string `json:"email" mapstructure:"email"`
	Role  string `json:"role,omitempty" mapstructure:"role"`
}

// JwtKey key to encode/decode JWT
var JwtKey []byte
