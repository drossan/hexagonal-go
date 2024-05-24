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

func TestMenuTreeHandler_CreateOrUpdateMenuTree_Integration(t *testing.T) {
	e := echo.New()
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	MenuTreeRepo := db.NewMenuTreeRepository(database)
	MenuTreeUseCase := usecase.NewMenuTreeUseCase(MenuTreeRepo)
	handler := api.NewMenuTreeHandler(e, MenuTreeUseCase)

	mockMenuTree := &model.MenuTree{Title: "Test Menu"}
	MenuTreeJSON, _ := json.Marshal(mockMenuTree)
	req := httptest.NewRequest(http.MethodPost, "/expanses-menus", bytes.NewBuffer(MenuTreeJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CreateOrUpdateExpanseMenu(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualResponse model.MenuTree
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockMenuTree.Title, actualResponse.Title)
	}
}

func TestMenuTreeHandler_GetAllMenuTree_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	MenuTreeRepo := db.NewMenuTreeRepository(database)
	MenuTreeUseCase := usecase.NewMenuTreeUseCase(MenuTreeRepo)
	MenuTreeHandler := api.NewMenuTreeHandler(e, MenuTreeUseCase)

	// Crear datos iniciales
	mockMenus := []*model.MenuTree{
		{Title: "Menu1"},
		{Title: "Menu2"},
	}
	for _, MenuTree := range mockMenus {
		database.Create(&MenuTree)
	}

	req := httptest.NewRequest(http.MethodGet, "/expanses-menus", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, MenuTreeHandler.GetAllExpanseMenus(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse []model.MenuTree
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Len(t, actualResponse, 2)
	}
}

func TestMenuTreeHandler_PaginateMenuTrees_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	MenuTreeRepo := db.NewMenuTreeRepository(database)
	MenuTreeUseCase := usecase.NewMenuTreeUseCase(MenuTreeRepo)
	MenuTreeHandler := api.NewMenuTreeHandler(e, MenuTreeUseCase)

	// Crear datos iniciales
	mockMenuTrees := []model.MenuTree{
		{
			Title: "Usuarios",
			Icon:  "mdi-account-check-outline",
			Order: 1,
		},
		{
			Title: "Usuarios 2",
			Icon:  "mdi-account-check-outline",
			Order: 2,
		},
	}

	for _, mockMenuTree := range mockMenuTrees {
		result := database.Create(&mockMenuTree)
		if result.Error != nil {
			t.Fatalf("Error creating mock data: %v", result.Error)
		}
	}

	req := httptest.NewRequest(http.MethodGet, "/expanses-menus/1?rows=", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/expanses-menus/:page")
	c.SetParamNames("page")
	c.SetParamValues("1")

	if assert.NoError(t, MenuTreeHandler.PaginateExpanseMenus(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, float64(2), actualResponse["total"].(float64))
		assert.Len(t, actualResponse["items"], 2)
	}
}

func TestMenuTreeHandler_DeleteMenuTree_Integration(t *testing.T) {
	e := echo.New()

	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)

	MenuTreeRepo := db.NewMenuTreeRepository(database)
	MenuTreeUseCase := usecase.NewMenuTreeUseCase(MenuTreeRepo)
	MenuTreeHandler := api.NewMenuTreeHandler(e, MenuTreeUseCase)

	// Crear un item inicial
	mockMenu := &model.MenuTree{Title: "Test Menu"}
	database.Create(&mockMenu)

	MenuTreeJSON, _ := json.Marshal(mockMenu)
	req := httptest.NewRequest(http.MethodPost, "/expanses-menus/delete", bytes.NewBuffer(MenuTreeJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, MenuTreeHandler.DeleteExpanseMenu(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.MenuTree
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockMenu.ID, actualResponse.ID)

		// Verificar que el item ha sido eliminado
		var count int64
		database.Model(&model.MenuTree{}).Where("id = ?", mockMenu.ID).Count(&count)
		assert.Equal(t, int64(0), count)
	}
}
