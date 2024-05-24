package integration_tests_test

import (
	"bytes"
	"encoding/json"
	"github.com/drossan/core-api/infrastructure/db"
	"github.com/drossan/core-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/interfaces/api"
	"github.com/drossan/core-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLevelPrivilegesHandler_CreateOrUpdateLevelPrivileges_Integration(t *testing.T) {
	e := echo.New()
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	LevelPrivilegesRepo := db.NewLevelPrivilegesRepository(database)
	LevelPrivilegesUseCase := usecase.NewLevelPrivilegesUseCase(LevelPrivilegesRepo)
	handler := api.NewLevelPrivilegesHandler(e, LevelPrivilegesUseCase)

	mockLevelPrivilege := &model.LevelPrivileges{FormID: 1, Read: true, Write: true}
	LevelPrivilegesJSON, _ := json.Marshal(mockLevelPrivilege)
	req := httptest.NewRequest(http.MethodPost, "/level-privilege", bytes.NewBuffer(LevelPrivilegesJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CreateLevelPrivilege(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualResponse model.LevelPrivileges
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockLevelPrivilege.FormID, actualResponse.FormID)
		assert.Equal(t, mockLevelPrivilege.Read, actualResponse.Read)
		assert.Equal(t, mockLevelPrivilege.Write, actualResponse.Write)
	}
}

func TestLevelPrivilegesHandler_GetAllLevelPrivileges_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	LevelPrivilegesRepo := db.NewLevelPrivilegesRepository(database)
	LevelPrivilegesUseCase := usecase.NewLevelPrivilegesUseCase(LevelPrivilegesRepo)
	LevelPrivilegesHandler := api.NewLevelPrivilegesHandler(e, LevelPrivilegesUseCase)

	// Crear datos iniciales
	mockLevelPrivileges := []*model.LevelPrivileges{
		{FormID: 1, Read: true, Write: true},
		{FormID: 2, Read: true, Write: false},
	}
	for _, LevelPrivileges := range mockLevelPrivileges {
		database.Create(&LevelPrivileges)
	}

	req := httptest.NewRequest(http.MethodGet, "/level-privileges", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, LevelPrivilegesHandler.GetAllLevelPrivileges(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse []model.LevelPrivileges
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Len(t, actualResponse, 2)
	}
}

func TestLevelPrivilegesHandler_DeleteLevelPrivileges_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	LevelPrivilegesRepo := db.NewLevelPrivilegesRepository(database)
	LevelPrivilegesUseCase := usecase.NewLevelPrivilegesUseCase(LevelPrivilegesRepo)
	LevelPrivilegesHandler := api.NewLevelPrivilegesHandler(e, LevelPrivilegesUseCase)

	// Crear un item inicial
	mockLevelPrivilege := &model.LevelPrivileges{FormID: 1, Read: true, Write: true}
	database.Create(&mockLevelPrivilege)

	LevelPrivilegesJSON, _ := json.Marshal(mockLevelPrivilege)
	req := httptest.NewRequest(http.MethodPost, "/level-privilege/delete", bytes.NewBuffer(LevelPrivilegesJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, LevelPrivilegesHandler.DeleteLevelPrivilege(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.LevelPrivileges
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockLevelPrivilege.ID, actualResponse.ID)

		// Verificar que el item ha sido eliminado
		var count int64
		database.Model(&model.LevelPrivileges{}).Where("id = ?", mockLevelPrivilege.ID).Count(&count)
		assert.Equal(t, int64(0), count)
	}
}
