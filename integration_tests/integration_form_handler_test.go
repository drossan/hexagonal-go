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

func TestFormHandler_CreateOrUpdateForm_Integration(t *testing.T) {
	e := echo.New()
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	formRepo := db.NewFormRepository(database)
	formUseCase := usecase.NewFormUseCase(formRepo)
	handler := api.NewFormHandler(e, formUseCase)

	mockForm := &model.Form{
		Title:   "Usuarios",
		Icon:    "mdi-account-check-outline",
		Link:    "usuarios",
		Setting: true,
		PathAPI: "user|users",
		Order:   1,
	}
	formJSON, _ := json.Marshal(mockForm)
	req := httptest.NewRequest(http.MethodPost, "/form", bytes.NewBuffer(formJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CreateOrUpdateForm(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualResponse model.Form
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockForm.Title, actualResponse.Title)
		assert.Equal(t, mockForm.Icon, actualResponse.Icon)
		assert.Equal(t, mockForm.Link, actualResponse.Link)
		assert.Equal(t, mockForm.Setting, actualResponse.Setting)
		assert.Equal(t, mockForm.PathAPI, actualResponse.PathAPI)
		assert.Equal(t, mockForm.Order, actualResponse.Order)
	}
}

func TestFormHandler_GetAllForm_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	formRepo := db.NewFormRepository(database)
	formUseCase := usecase.NewFormUseCase(formRepo)
	formHandler := api.NewFormHandler(e, formUseCase)

	// Crear datos iniciales
	forms := []model.Form{
		{Title: "Form 1", Icon: "icon-1", Link: "link-1", Setting: true, PathAPI: "path-1", Order: 1},
		{Title: "Form 2", Icon: "icon-2", Link: "link-2", Setting: true, PathAPI: "path-2", Order: 2},
	}
	for _, form := range forms {
		database.Create(&form)
	}

	req := httptest.NewRequest(http.MethodGet, "/forms", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, formHandler.GetAllForms(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse []model.Form
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Len(t, actualResponse, 2)
	}
}

func TestFormHandler_PaginateForms_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	formRepo := db.NewFormRepository(database)
	formUseCase := usecase.NewFormUseCase(formRepo)
	formHandler := api.NewFormHandler(e, formUseCase)

	// Crear datos iniciales
	forms := []model.Form{
		{Title: "Form 1", Icon: "icon-1", Link: "link-1", Setting: true, PathAPI: "path-1", Order: 1},
		{Title: "Form 2", Icon: "icon-2", Link: "link-2", Setting: true, PathAPI: "path-2", Order: 2},
		{Title: "Form 3", Icon: "icon-3", Link: "link-3", Setting: true, PathAPI: "path-3", Order: 3},
	}
	for _, form := range forms {
		database.Create(&form)
	}

	req := httptest.NewRequest(http.MethodGet, "/forms/1?rows=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/forms/:page")
	c.SetParamNames("page")
	c.SetParamValues("1")

	if assert.NoError(t, formHandler.PaginateForms(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, float64(3), actualResponse["total"].(float64))
		assert.Len(t, actualResponse["items"], 2)
	}
}

func TestFormHandler_DeleteForm_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	formRepo := db.NewFormRepository(database)
	formUseCase := usecase.NewFormUseCase(formRepo)
	formHandler := api.NewFormHandler(e, formUseCase)

	// Crear un formulario inicial
	form := model.Form{Title: "Form to delete", Icon: "icon", Link: "link", Setting: true, PathAPI: "path", Order: 1}
	database.Create(&form)

	formJSON, _ := json.Marshal(form)
	req := httptest.NewRequest(http.MethodPost, "/form/delete", bytes.NewBuffer(formJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, formHandler.DeleteForm(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.Form
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, form.ID, actualResponse.ID)

		// Verificar que el formulario ha sido eliminado
		var count int64
		database.Model(&model.Form{}).Where("id = ?", form.ID).Count(&count)
		assert.Equal(t, int64(0), count)
	}
}
