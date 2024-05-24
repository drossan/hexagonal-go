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

func TestMenuTreeHandler_CreateOrUpdateMenuTree(t *testing.T) {
	e := echo.New()
	mockRepo := new(mocks.MockMenuTreeRepository)

	MenuTreeUseCase := usecase.NewMenuTreeUseCase(mockRepo)
	handler := api.NewMenuTreeHandler(e, MenuTreeUseCase)

	mockMenuTree := &model.MenuTree{
		Title: "Usuarios",
		Icon:  "mdi-account-check-outline",
		Order: 1,
	}
	MenuTreeJSON, _ := json.Marshal(mockMenuTree)
	req := httptest.NewRequest(http.MethodPost, "/expanses-menus", bytes.NewBuffer(MenuTreeJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo.On("CreateOrUpdate", mock.Anything).Return(nil)

	if assert.NoError(t, handler.CreateOrUpdateExpanseMenu(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualResponse model.MenuTree
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockMenuTree.Title, actualResponse.Title)
		assert.Equal(t, mockMenuTree.Icon, actualResponse.Icon)
		assert.Equal(t, mockMenuTree.Order, actualResponse.Order)
	}
}

func TestMenuTreeHandler_PaginateMenuTrees(t *testing.T) {
	e := echo.New()
	mockRepo := new(mocks.MockMenuTreeRepository)
	mockMenuTree := []*model.MenuTree{
		{
			Model: gorm.Model{ID: 1},
			Title: "Usuarios",
			Icon:  "mdi-account-check-outline",
			Order: 1,
		},
		{
			Model: gorm.Model{ID: 2},
			Title: "Identidades",
			Icon:  "mdi-account-check-outline",
			Order: 2,
		},
	}

	mockRepo.On("Paginate", 1, 2).Return(mockMenuTree, 2, nil)

	MenuTreeUseCase := usecase.NewMenuTreeUseCase(mockRepo)
	handler := api.NewMenuTreeHandler(e, MenuTreeUseCase)

	req := httptest.NewRequest(http.MethodGet, "/expanses-menus/1?rows=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/MenuTrees/:page")
	c.SetParamNames("page")
	c.SetParamValues("1")

	if assert.NoError(t, handler.PaginateExpanseMenus(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, float64(2), actualResponse["total"].(float64))
		assert.Len(t, actualResponse["items"], 2)
	}
}

func TestMenuTreeHandler_DeleteMenuTree(t *testing.T) {
	e := echo.New()

	mockRepo := new(mocks.MockMenuTreeRepository)

	MenuTreeUseCase := usecase.NewMenuTreeUseCase(mockRepo)
	handler := api.NewMenuTreeHandler(e, MenuTreeUseCase)

	mockMenuTree := &model.MenuTree{
		Model: gorm.Model{ID: 1},
	}
	MenuTreeJSON, _ := json.Marshal(mockMenuTree)
	req := httptest.NewRequest(http.MethodPost, "/expanses-menus/delete", bytes.NewBuffer(MenuTreeJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo.On("Delete", mock.Anything).Return(nil)

	if assert.NoError(t, handler.DeleteExpanseMenu(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualResponse model.MenuTree
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockMenuTree.ID, actualResponse.ID)
	}
}
