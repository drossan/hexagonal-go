package integration_tests_test

import (
	"bytes"
	"encoding/json"
	"github.com/drossan/core-api/infrastructure/db"
	"github.com/drossan/core-api/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/interfaces/api"
	"github.com/drossan/core-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_CreateOrUpdateUser_Integration(t *testing.T) {
	e := echo.New()
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	UserRepo := db.NewUserRepository(database)
	UserUseCase := usecase.NewUserUseCase(UserRepo)
	handler := api.NewUserHandler(e, UserUseCase)

	mockUser := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		FullName: "Test User",
		Password: "password",
	}

	UserJSON, _ := json.Marshal(mockUser)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(UserJSON))
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
		assert.Equal(t, mockUser.FullName, actualResponse.FullName)

		// Verificar la contraseña hasheada
		err = bcrypt.CompareHashAndPassword([]byte(actualResponse.Password), []byte(mockUser.Password))
		assert.NoError(t, err, "The password should be hashed correctly and match the original password")
	}
}

func TestUserHandler_PaginateUsers_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	UserRepo := db.NewUserRepository(database)
	UserUseCase := usecase.NewUserUseCase(UserRepo)
	UserHandler := api.NewUserHandler(e, UserUseCase)

	// Crear datos iniciales
	Users := []model.User{
		{
			Username: "testuser1",
			Email:    "test1@example.com",
			FullName: "Test User 1",
			Password: "password1",
		},
		{
			Username: "testuser2",
			Email:    "test2@example.com",
			FullName: "Test User 2",
			Password: "password2",
		},
		{
			Username: "testuser3",
			Email:    "test3@example.com",
			FullName: "Test User 3",
			Password: "password3",
		},
	}
	for _, User := range Users {
		database.Create(&User)
	}

	req := httptest.NewRequest(http.MethodGet, "/users/1?rows=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("s/:page")
	c.SetParamNames("page")
	c.SetParamValues("1")

	if assert.NoError(t, UserHandler.PaginateUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, float64(3), actualResponse["total"].(float64))
		assert.Len(t, actualResponse["items"], 2)
	}
}

func TestUserHandler_GetUserData_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	UserRepo := db.NewUserRepository(database)
	UserUseCase := usecase.NewUserUseCase(UserRepo)
	UserHandler := api.NewUserHandler(e, UserUseCase)

	// Crear un item inicial
	mockUser := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		FullName: "Test User",
		Password: "password",
	}
	database.Create(&mockUser)

	// Crear un token JWT para el usuario
	claims := &model.Claim{
		UserID:           mockUser.ID,
		Email:            mockUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	UserJSON, _ := json.Marshal(mockUser)
	req := httptest.NewRequest(http.MethodGet, "/user", bytes.NewBuffer(UserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+tokenString)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Agregar el token al contexto
	c.Set("user", token)

	if assert.NoError(t, UserHandler.GetUserData(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.User
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, actualResponse.ID)
		assert.Equal(t, mockUser.Username, actualResponse.Username)
		assert.Equal(t, mockUser.Email, actualResponse.Email)
		assert.Equal(t, mockUser.FullName, actualResponse.FullName)
		assert.Equal(t, mockUser.Password, actualResponse.Password)
	}
}

func TestUserHandler_DeleteUser_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	UserRepo := db.NewUserRepository(database)
	UserUseCase := usecase.NewUserUseCase(UserRepo)
	UserHandler := api.NewUserHandler(e, UserUseCase)

	// Crear un item inicial
	mockUser := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		FullName: "Test User",
		Password: "password",
	}
	database.Create(&mockUser)

	UserJSON, _ := json.Marshal(mockUser)
	req := httptest.NewRequest(http.MethodPost, "/user/delete", bytes.NewBuffer(UserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, UserHandler.DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.User
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, actualResponse.ID)

		// Verificar que el item ha sido eliminado
		var count int64
		database.Model(&model.User{}).Where("id = ?", mockUser.ID).Count(&count)
		assert.Equal(t, int64(0), count)
	}
}
