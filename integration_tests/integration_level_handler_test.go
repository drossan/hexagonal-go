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

func TestLevelHandler_CreateOrUpdateLevel_Integration(t *testing.T) {
	e := echo.New()
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	LevelRepo := db.NewLevelRepository(database)
	LevelUseCase := usecase.NewLevelUseCase(LevelRepo)
	handler := api.NewLevelHandler(e, LevelUseCase)

	mockLevel := &model.Level{Level: "Test Level", Description: "Level mock 1"}
	LevelJSON, _ := json.Marshal(mockLevel)
	req := httptest.NewRequest(http.MethodPost, "/level", bytes.NewBuffer(LevelJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CreateOrUpdateLevel(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualResponse model.Level
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockLevel.Level, actualResponse.Level)
	}
}

func TestLevelHandler_GetAllLevel_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	LevelRepo := db.NewLevelRepository(database)
	LevelUseCase := usecase.NewLevelUseCase(LevelRepo)
	LevelHandler := api.NewLevelHandler(e, LevelUseCase)

	// Crear datos iniciales
	mockLevels := []*model.Level{
		{Level: "Level1", Description: "Level mock 1"},
		{Level: "Level2", Description: "Level mock 2"},
		{Level: "Level3", Description: "Level mock 3"},
	}

	for _, mockLevel := range mockLevels {
		database.Create(&mockLevel)
	}

	req := httptest.NewRequest(http.MethodGet, "/levels", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, LevelHandler.GetAllLevels(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse []model.Level
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Len(t, actualResponse, 3)
	}
}

func TestLevelHandler_PaginateLevels_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	LevelRepo := db.NewLevelRepository(database)
	LevelUseCase := usecase.NewLevelUseCase(LevelRepo)
	LevelHandler := api.NewLevelHandler(e, LevelUseCase)

	// Crear datos iniciales
	mockLevels := []*model.Level{
		{Level: "Level1", Description: "Level mock 1"},
		{Level: "Level2", Description: "Level mock 2"},
		{Level: "Level3", Description: "Level mock 3"},
	}

	for _, mockLevel := range mockLevels {
		if err := database.Create(&mockLevel).Error; err != nil {
			t.Fatalf("Error inserting level: %v", err)
		}
	}

	// Verificar que los niveles se insertaron correctamente
	var count int64
	if err := database.Model(&model.Level{}).Count(&count).Error; err != nil {
		t.Fatalf("Error counting levels: %v", err)
	}
	if count != 3 {
		t.Fatalf("Expected 3 levels in the database, but found %d", count)
	}

	req := httptest.NewRequest(http.MethodGet, "/levels/1?rows=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/levels/:page")
	c.SetParamNames("page")
	c.SetParamValues("1")

	if assert.NoError(t, LevelHandler.PaginateLevels(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, float64(3), actualResponse["total"].(float64))
		assert.Len(t, actualResponse["items"], 2) // Debe devolver solo 2 ítems debido a la paginación

		// Verificar los ítems devueltos
		items, ok := actualResponse["items"].([]interface{})
		if !ok {
			t.Fatalf("Expected items to be an array")
		}
		for i, item := range items {
			levelMap := item.(map[string]interface{})
			expectedLevel := mockLevels[i]
			assert.Equal(t, expectedLevel.Level, levelMap["level"])
			assert.Equal(t, expectedLevel.Description, levelMap["description"])
		}
	}
}

func TestLevelHandler_DeleteLevel_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	LevelRepo := db.NewLevelRepository(database)
	LevelUseCase := usecase.NewLevelUseCase(LevelRepo)
	LevelHandler := api.NewLevelHandler(e, LevelUseCase)

	// Crear un item inicial
	mockLevel := &model.Level{Level: "Test Level", Description: "Level mock 1"}
	database.Create(&mockLevel)

	LevelJSON, _ := json.Marshal(mockLevel)
	req := httptest.NewRequest(http.MethodPost, "/levels/delete", bytes.NewBuffer(LevelJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, LevelHandler.DeleteLevel(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.Level
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockLevel.ID, actualResponse.ID)

		// Verificar que el item ha sido eliminado
		var count int64
		database.Model(&model.Level{}).Where("id = ?", mockLevel.ID).Count(&count)
		assert.Equal(t, int64(0), count)
	}
}
