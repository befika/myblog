package model

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type UserClaims struct {
	jwt.StandardClaims
	Email  string    `json:"email"`
	UserId uuid.UUID `json:"user_id"`
	Roles  []Role    `json:"roles"`
}
type LoginResponse struct {
	Token  string      `json:"token"`
	Name   string      `json:"name"`
	Email  string      `json:"email"`
	UserId uuid.UUID   `json:"user_id"`
	Roles  []UserRoles `json:"roles"`
}
