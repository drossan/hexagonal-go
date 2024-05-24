package middleware

import (
	"net/http"
	"strings"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// AuthorizationMiddlewareConfig guarda las dependencias necesarias para el middleware
type AuthorizationMiddlewareConfig struct {
	LevelRepo           repository.LevelRepository
	FormRepo            repository.FormRepository
	LevelPrivilegesRepo repository.LevelPrivilegesRepository
}

// AuthorizationMiddleware verifica los permisos del usuario
func AuthorizationMiddleware(config AuthorizationMiddlewareConfig, prefix string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userToken := c.Get("user").(*jwt.Token)
			claims := userToken.Claims.(*model.Claim)
			levelID := claims.LevelID

			// Obtener el path después de /api/v1/
			path := strings.TrimPrefix(c.Path(), "/"+prefix+"/")
			segments := strings.Split(path, "/")
			path = segments[0]

			// Obtener el nivel del usuario
			level, err := config.LevelRepo.GetByID(levelID)
			if err != nil {
				return c.JSON(http.StatusForbidden, map[string]interface{}{"error": "01 - Access denied"})
			}

			// Verificar los privilegios del nivel del usuario
			hasAccess := false
			method := c.Request().Method

			for _, privilege := range level.LevelPrivileges {
				pathAPI := strings.Split(privilege.Form.PathAPI, "|")

				if len(pathAPI) > 1 && (pathAPI[0] == path || pathAPI[1] == path) {
					if method == http.MethodGet && privilege.Read {
						hasAccess = true
						break
					} else if (method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete) && privilege.Write {
						hasAccess = true
						break
					}
				}
			}

			if !hasAccess {
				return c.JSON(http.StatusForbidden, map[string]interface{}{"error": "02 - Access denied"})
			}

			return next(c)
		}
	}
}

// NewAuthorizationMiddleware crea una nueva instancia de AuthorizationMiddlewareConfig y la devuelve como un middleware
func NewAuthorizationMiddleware(levelRepo repository.LevelRepository, formRepo repository.FormRepository, levelPrivilegesRepo repository.LevelPrivilegesRepository, prefix string) echo.MiddlewareFunc {
	config := AuthorizationMiddlewareConfig{
		LevelRepo:           levelRepo,
		FormRepo:            formRepo,
		LevelPrivilegesRepo: levelPrivilegesRepo,
	}

	return AuthorizationMiddleware(config, prefix)
}
