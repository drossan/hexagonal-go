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
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFormHandler_CreateOrUpdateForm(t *testing.T) {
	e := echo.New()
	mockRepo := new(mocks.MockFormRepository)

	FormUseCase := usecase.NewFormUseCase(mockRepo)
	handler := api.NewFormHandler(e, FormUseCase)

	mockForm := &model.Form{
		Title:   "Usuarios",
		Icon:    "mdi-account-check-outline",
		Link:    "usuarios",
		Setting: true,
		PathAPI: "user|users",
		Order:   1,
	}
	FormJSON, _ := json.Marshal(mockForm)
	req := httptest.NewRequest(http.MethodPost, "/form", bytes.NewBuffer(FormJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Definir la expectativa de la llamada al método CreateOrUpdate
	mockRepo.On("CreateOrUpdate", mock.Anything).Return(nil)

	if assert.NoError(t, handler.CreateOrUpdateForm(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualResponse model.Form
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockForm.Title, actualResponse.Title)
		assert.Equal(t, mockForm.Icon, actualResponse.Icon)
		assert.Equal(t, mockForm.Link, actualResponse.Link)
		assert.Equal(t, mockForm.Link, actualResponse.Link)
		assert.Equal(t, mockForm.Setting, actualResponse.Setting)
		assert.Equal(t, mockForm.PathAPI, actualResponse.PathAPI)
		assert.Equal(t, mockForm.Order, actualResponse.Order)
	}

	// Verificar que se cumplieron las expectativas
	mockRepo.AssertExpectations(t)
}

func TestFormHandler_PaginateForms(t *testing.T) {
	e := echo.New()
	mockRepo := new(mocks.MockFormRepository)
	mockForms := []*model.Form{
		{
			Model:   gorm.Model{ID: 1},
			Title:   "Usuarios",
			Icon:    "mdi-account-check-outline",
			Link:    "usuarios",
			Setting: true,
			PathAPI: "user|users",
			Order:   1,
		},
		{
			Model:   gorm.Model{ID: 2},
			Title:   "Identidades",
			Icon:    "mdi-account-check-outline",
			Link:    "roles",
			Setting: true,
			PathAPI: "level|levels",
			Order:   2,
		},
	}
	mockRepo.On("Paginate", 1, 2).Return(mockForms, 2, nil)

	FormUseCase := usecase.NewFormUseCase(mockRepo)
	handler := api.NewFormHandler(e, FormUseCase)

	req := httptest.NewRequest(http.MethodGet, "/forms/1?rows=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/Forms/:page")
	c.SetParamNames("page")
	c.SetParamValues("1")

	if assert.NoError(t, handler.PaginateForms(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, float64(2), actualResponse["total"].(float64))
		assert.Len(t, actualResponse["items"], 2)
	}
}

func TestFormHandler_DeleteForm(t *testing.T) {
	e := echo.New()

	mockRepo := new(mocks.MockFormRepository)

	FormUseCase := usecase.NewFormUseCase(mockRepo)
	handler := api.NewFormHandler(e, FormUseCase)

	mockForm := &model.Form{
		Model: gorm.Model{ID: 1},
	}

	mockRepo.On("Delete", mock.Anything).Return(nil)

	FormJSON, _ := json.Marshal(mockForm)
	req := httptest.NewRequest(http.MethodPost, "/form/delete", bytes.NewBuffer(FormJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.DeleteForm(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.Form
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockForm.ID, actualResponse.ID)
	}
}
