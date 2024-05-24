package router

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func accessible(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>API</h1>")
}

func NewEchoRouter(JWTSecret string) (*echo.Echo, *echo.Group, *echo.Group, string) {
	e := echo.New()

	prefix := "api/v1"

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Public path
	e.Static("/", "public")

	// Unauthenticated route
	e.GET(prefix, accessible)

	// Accessible group
	a := e.Group(prefix)

	// Restricted group
	r := e.Group(prefix)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.Claim)
		},
		SigningKey: []byte(JWTSecret),
	}
	r.Use(echojwt.WithConfig(config))

	return e, r, a, prefix
}
