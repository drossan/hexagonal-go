package model

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claim Token de user
type Claim struct {
	UserID  uint   `json:"user_id"`
	Email   string `json:"email"`
	LevelID uint   `json:"level_id"`
	Token   string `json:"token"`
	Admin   uint
	jwt.RegisteredClaims
}
