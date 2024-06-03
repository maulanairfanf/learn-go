package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims represents the JWT claims
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}
