package helpers

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

func GetCurrentUser(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.Claim)
	return claims.UserID
}

func GenerateJWT(user *model.User) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		Issuer:    "Intranet API - Generate by IslaIT",
	}

	claims := model.Claim{
		UserID:           user.ID,
		Email:            user.Email,
		Token:            user.Token,
		LevelID:          user.LevelID,
		Admin:            user.LevelID,
		RegisteredClaims: registeredClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
