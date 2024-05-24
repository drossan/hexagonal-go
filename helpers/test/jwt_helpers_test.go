package helpers_test

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/helpers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// Test unitario para GetCurrentUser
func TestGetCurrentUser(t *testing.T) {
	// Crear un nuevo contexto Echo
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Crear las reclamaciones y el token JWT
	claims := &model.Claim{
		UserID: 123,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Configurar el contexto Echo con el token JWT
	c.Set("user", token)

	// Llamar a la función y verificar el resultado
	userID := helpers.GetCurrentUser(c)
	assert.Equal(t, uint(123), userID)
}

func TestGenerateJWT(t *testing.T) {
	// Simular la variable de entorno JWT_SECRET
	_ = os.Setenv("JWT_SECRET", "test_secret")
	defer func() {
		_ = os.Unsetenv("JWT_SECRET")
	}()

	// Crear un usuario de prueba
	user := &model.User{
		Model:   gorm.Model{ID: 1},
		Email:   "test@example.com",
		Token:   "some_token",
		LevelID: 1,
	}

	// Generar el token JWT
	tokenString, err := helpers.GenerateJWT(user)

	// Verificar que no hubo errores
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Verificar el contenido del token JWT
	token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("test_secret"), nil
	})

	assert.NoError(t, err)
	assert.NotNil(t, token)

	claims, ok := token.Claims.(*model.Claim)
	assert.True(t, ok)
	assert.Equal(t, user.ID, claims.UserID)
	assert.Equal(t, user.Email, claims.Email)
	assert.Equal(t, user.Token, claims.Token)
	assert.Equal(t, user.LevelID, claims.LevelID)
	assert.Equal(t, user.LevelID, claims.Admin)
	assert.Equal(t, "Intranet API - Generate by IslaIT", claims.Issuer)
	assert.WithinDuration(t, time.Now().Add(time.Hour*72), claims.ExpiresAt.Time, time.Minute)
}
