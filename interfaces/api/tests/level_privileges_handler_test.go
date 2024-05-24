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

func TestLevelPrivilegesHandler_CreateOrUpdateLevel(t *testing.T) {
	e := echo.New()

	mockRepo := new(mocks.MockLevelPrivilegesRepository)

	LevelUseCase := usecase.NewLevelPrivilegesUseCase(mockRepo)
	handler := api.NewLevelPrivilegesHandler(e, LevelUseCase)

	mockLevel := &model.LevelPrivileges{FormID: 1, Read: true, Write: true}
	LevelJSON, _ := json.Marshal(mockLevel)
	req := httptest.NewRequest(http.MethodPost, "/level-privilege", bytes.NewBuffer(LevelJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo.On("CreateOrUpdate", mock.Anything).Return(nil)

	if assert.NoError(t, handler.CreateLevelPrivilege(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualResponse model.LevelPrivileges
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockLevel.FormID, actualResponse.FormID)
	}
}

func TestLevelPrivilegesHandler_DeleteLevel(t *testing.T) {
	e := echo.New()

	mockRepo := new(mocks.MockLevelPrivilegesRepository)

	LevelUseCase := usecase.NewLevelPrivilegesUseCase(mockRepo)
	handler := api.NewLevelPrivilegesHandler(e, LevelUseCase)

	mockLevel := &model.LevelPrivileges{FormID: 1, Read: true, Write: true}

	LevelJSON, _ := json.Marshal(mockLevel)
	req := httptest.NewRequest(http.MethodPost, "/level-privilege/delete", bytes.NewBuffer(LevelJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo.On("Delete", mock.Anything).Return(nil)

	if assert.NoError(t, handler.DeleteLevelPrivilege(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.LevelPrivileges
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockLevel.ID, actualResponse.ID)
	}
}
