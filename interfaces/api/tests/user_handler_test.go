package api_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/interfaces/api"
	"github.com/drossan/core-api/interfaces/api/tests/mocks"
	"github.com/drossan/core-api/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// generateToken is a helper function to generate a JWT token for testing
func generateToken(userID uint, jwtSecret string) (string, error) {
	claims := model.Claim{
		UserID:  userID,
		Email:   "test@example.com",
		LevelID: 1,
		Admin:   1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			Issuer:    "Intranet API - generated by IslaIT",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func TestUserHandler_CreateOrUpdateUser(t *testing.T) {
	e := echo.New()

	mockUserRepo := &mocks.MockUserRepository{
		CreateFunc: func(user *model.User) error {
			user.ID = 1
			return nil
		},
	}

	userUseCase := usecase.NewUserUseCase(mockUserRepo)
	handler := api.NewUserHandler(e, userUseCase)

	mockUser := &model.User{Username: "testuser", Email: "test@example.com"}
	userJSON, _ := json.Marshal(mockUser)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CreateOrUpdateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualResponse model.User
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.Username, actualResponse.Username)
		assert.Equal(t, mockUser.Email, actualResponse.Email)
	}
}

func TestUserHandler_GetUserData(t *testing.T) {
	e := echo.New()

	mockUserRepo := &mocks.MockUserRepository{
		GetByIDFunc: func(id uint) (*model.User, error) {
			return &model.User{
				Model:    gorm.Model{ID: id},
				Username: "testuser",
				Email:    "test@example.com",
			}, nil
		},
	}

	userUseCase := usecase.NewUserUseCase(mockUserRepo)
	handler := api.NewUserHandler(e, userUseCase)

	token, err := generateToken(1, "testsecret")
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &jwt.Token{
		Claims: &model.Claim{
			UserID: 1,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		},
	})

	if assert.NoError(t, handler.GetUserData(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.User
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), actualResponse.ID)
		assert.Equal(t, "testuser", actualResponse.Username)
		assert.Equal(t, "test@example.com", actualResponse.Email)
	}
}

func TestUserHandler_PaginateUsers(t *testing.T) {
	e := echo.New()

	mockUserRepo := &mocks.MockUserRepository{
		PaginateFunc: func(page int, pageSize int) ([]*model.User, int, error) {
			users := []*model.User{
				{
					Model:    gorm.Model{ID: 1},
					Username: "testuser1",
					Email:    "test1@example.com",
				},
				{
					Model:    gorm.Model{ID: 2},
					Username: "testuser2",
					Email:    "test2@example.com",
				},
			}
			return users, 2, nil
		},
	}

	userUseCase := usecase.NewUserUseCase(mockUserRepo)
	handler := api.NewUserHandler(e, userUseCase)

	req := httptest.NewRequest(http.MethodGet, "/users/1?rows=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:page")
	c.SetParamNames("page")
	c.SetParamValues("1")

	if assert.NoError(t, handler.PaginateUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, float64(2), actualResponse["total"].(float64))
		assert.Len(t, actualResponse["items"], 2)
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	e := echo.New()

	mockUserRepo := &mocks.MockUserRepository{
		DeleteFunc: func(user *model.User) error {
			return nil
		},
	}

	userUseCase := usecase.NewUserUseCase(mockUserRepo)
	handler := api.NewUserHandler(e, userUseCase)

	mockUser := &model.User{
		Model: gorm.Model{ID: 1},
	}
	userJSON, _ := json.Marshal(mockUser)
	req := httptest.NewRequest(http.MethodPost, "/user/delete", bytes.NewBuffer(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.User
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, actualResponse.ID)
	}
}

func TestUserHandler_Login(t *testing.T) {
	e := echo.New()

	// Crear los mocks
	mockUserRepo := &mocks.MockUserRepository{
		GetByEmailFunc: func(email string) (*model.User, error) {
			if email == "test@example.com" {
				ps := sha256.Sum256([]byte("password"))
				pwd := fmt.Sprintf("%x", ps)
				return &model.User{
					Model:    gorm.Model{ID: 1},
					Email:    "test@example.com",
					Password: pwd,
				}, nil
			}
			return nil, errors.New("invalid email or password")
		},
		LoginFunc: func(email, password string) (string, error) {
			if email == "test@example.com" && password == "password" {
				return "mockToken", nil
			}
			return "", errors.New("invalid email or password")
		},
	}

	userUseCase := usecase.NewUserUseCase(mockUserRepo)
	handler := api.NewUserHandler(e, userUseCase)

	// Prueba de login exitoso
	mockCredentials := map[string]string{
		"email":    "test@example.com",
		"password": "password",
	}
	credentialsJSON, _ := json.Marshal(mockCredentials)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(credentialsJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		// Validar el token JWT
		tokenString := actualResponse["data"].(map[string]interface{})["token"].(string)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("test_secret"), nil
		})
		assert.NoError(t, err)
		claims, ok := token.Claims.(jwt.MapClaims)
		assert.True(t, ok)
		assert.Equal(t, float64(1), claims["user_id"])
		assert.Equal(t, "test@example.com", claims["email"])
		assert.Equal(t, "Intranet API - generate by IslaIT", claims["iss"])
	}

	// Prueba de login fallido
	mockCredentials = map[string]string{
		"email":    "test@example.com",
		"password": "wrongpassword",
	}
	credentialsJSON, _ = json.Marshal(mockCredentials)
	req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(credentialsJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		expectedResponse := map[string]interface{}{
			"error": "invalid email or password",
		}
		assert.Equal(t, expectedResponse, actualResponse)
	}
}
