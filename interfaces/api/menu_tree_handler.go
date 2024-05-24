package api

import (
	"net/http"
	"strconv"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/usecase"
	"github.com/labstack/echo/v4"
)

type MenuTreeHandler struct {
	menuTreeUseCase *usecase.MenuTreeUseCase
}

func NewMenuTreeHandler(e *echo.Echo, uc *usecase.MenuTreeUseCase) *MenuTreeHandler {
	return &MenuTreeHandler{menuTreeUseCase: uc}
}

func (h *MenuTreeHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/expanses-menus", h.GetAllExpanseMenus)
	g.GET("/expanses-menus/:page", h.PaginateExpanseMenus)
	g.POST("/expanses-menus", h.CreateOrUpdateExpanseMenu)
	g.POST("/expanses-menus/delete", h.DeleteExpanseMenu)
}

// GetAllExpanseMenus godoc
// @Summary Get all expanse menus
// @Description Get all expanse menus
// @Tags expanse menus
// @Accept json
// @Produce json
// @Success 200 {array} model.MenuTree
// @Failure 500 {object} map[string]interface{}
// @Router /expanses-menus [get]
func (h *MenuTreeHandler) GetAllExpanseMenus(c echo.Context) error {
	menus, err := h.menuTreeUseCase.GetAllMenuTrees()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, menus)
}

// PaginateExpanseMenus godoc
// @Summary Get expanse menus with pagination
// @Description Get expanse menus with pagination
// @Tags expanse menus
// @Accept json
// @Produce json
// @Param page path int true "Page number"
// @Param rows query int false "Rows per page"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /expanses-menus/{page} [get]
func (h *MenuTreeHandler) PaginateExpanseMenus(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil || page < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid page number"})
	}

	rows, err := strconv.Atoi(c.QueryParam("rows"))
	if err != nil || rows < 1 {
		rows = 50
	}

	menus, total, err := h.menuTreeUseCase.PaginateMenuTrees(page, rows)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": menus,
		"total": total,
	})
}

// CreateOrUpdateExpanseMenu godoc
// @Summary Add a new expanse menu
// @Description Add a new expanse menu
// @Tags expanse menus
// @Accept json
// @Produce json
// @Param menu body model.MenuTree true "Expanse Menu"
// @Success 201 {object} model.MenuTree
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /expanses-menus [post]
func (h *MenuTreeHandler) CreateOrUpdateExpanseMenu(c echo.Context) error {
	menu := new(model.MenuTree)
	if err := c.Bind(menu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := h.menuTreeUseCase.CreateOrUpdateMenuTree(menu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, menu)
}

// DeleteExpanseMenu godoc
// @Summary Delete an expanse menu
// @Description Delete an expanse menu by ID
// @Tags expanse menus
// @Accept json
// @Produce json
// @Param id query int true "Menu ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /expanses-menus/delete [post]
func (h *MenuTreeHandler) DeleteExpanseMenu(c echo.Context) error {

	menu := new(model.MenuTree)
	if err := c.Bind(menu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := h.menuTreeUseCase.DeleteMenuTree(menu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, menu)
}
