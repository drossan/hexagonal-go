package api_test

import (
	"bytes"
	"encoding/json"
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/interfaces/api"
	"github.com/drossan/core-api/mocks"
	"github.com/drossan/core-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLevelHandler_CreateOrUpdateLevel(t *testing.T) {
	e := echo.New()

	mockRepo := new(mocks.MockLevelRepository)

	LevelUseCase := usecase.NewLevelUseCase(mockRepo)
	handler := api.NewLevelHandler(e, LevelUseCase)

	mockLevel := &model.Level{Level: "Test Level"}
	LevelJSON, _ := json.Marshal(mockLevel)
	req := httptest.NewRequest(http.MethodPost, "/Level", bytes.NewBuffer(LevelJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo.On("CreateOrUpdate", mock.Anything).Return(nil)

	if assert.NoError(t, handler.CreateOrUpdateLevel(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualResponse model.Level
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockLevel.Level, actualResponse.Level)
	}
}

func TestLevelHandler_PaginateLevels(t *testing.T) {
	e := echo.New()
	mockRepo := new(mocks.MockLevelRepository)
	mockLevels := []*model.Level{
		{Level: "Level1"},
		{Level: "Level2"},
	}
	mockRepo.On("Paginate", 1, 2).Return(mockLevels, 2, nil)

	LevelUseCase := usecase.NewLevelUseCase(mockRepo)
	handler := api.NewLevelHandler(e, LevelUseCase)

	req := httptest.NewRequest(http.MethodGet, "/Levels/1?rows=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/Levels/:page")
	c.SetParamNames("page")
	c.SetParamValues("1")

	if assert.NoError(t, handler.PaginateLevels(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, float64(2), actualResponse["total"].(float64))
		assert.Len(t, actualResponse["items"], 2)
	}

	// Verificar que se cumplieron las expectativas
	mockRepo.AssertExpectations(t)
}

func TestLevelHandler_DeleteLevel(t *testing.T) {
	e := echo.New()

	mockRepo := new(mocks.MockLevelRepository)

	LevelUseCase := usecase.NewLevelUseCase(mockRepo)
	handler := api.NewLevelHandler(e, LevelUseCase)

	mockLevel := &model.Level{Level: "Test Level"}

	LevelJSON, _ := json.Marshal(mockLevel)
	req := httptest.NewRequest(http.MethodPost, "/Level/delete", bytes.NewBuffer(LevelJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo.On("Delete", mock.Anything).Return(nil)

	if assert.NoError(t, handler.DeleteLevel(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.Level
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockLevel.ID, actualResponse.ID)
	}
}
