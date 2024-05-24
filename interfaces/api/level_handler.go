package api

import (
	"net/http"
	"strconv"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/usecase"
	"github.com/labstack/echo/v4"
)

// LevelHandler manages levels
type LevelHandler struct {
	levelUseCase *usecase.LevelUseCase
}

// NewLevelHandler initializes a new LevelHandler
func NewLevelHandler(e *echo.Echo, uc *usecase.LevelUseCase) *LevelHandler {
	return &LevelHandler{levelUseCase: uc}
}

// RegisterRoutes registers level routes
func (h *LevelHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/levels", h.GetAllLevels)
	g.GET("/levels/:page", h.PaginateLevels)
	g.POST("/level", h.CreateOrUpdateLevel)
	g.POST("/level/delete", h.DeleteLevel)
}

// GetAllLevels godoc
// @Summary Get all levels
// @Description Get all levels
// @Tags levels
// @Accept json
// @Produce json
// @Success 200 {object} []model.Level
// @Failure 500 {object} map[string]interface{}
// @Router /levels [get]
func (h *LevelHandler) GetAllLevels(c echo.Context) error {
	levels, err := h.levelUseCase.GetAllLevels()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, levels)
}

// PaginateLevels godoc
// @Summary Get levels with pagination
// @Description Get levels with pagination
// @Tags levels
// @Accept json
// @Produce json
// @Param page path int true "Page number"
// @Param rows query int false "Rows per page"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /levels/{page} [get]
func (h *LevelHandler) PaginateLevels(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil || page < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid page number"})
	}

	rows, err := strconv.Atoi(c.QueryParam("rows"))
	if err != nil || rows < 1 {
		rows = 50
	}

	levels, total, err := h.levelUseCase.PaginateLevels(page, rows)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": levels,
		"total": total,
	})
}

// CreateOrUpdateLevel godoc
// @Summary Add a new level
// @Description Add a new level with the input payload
// @Tags levels
// @Accept json
// @Produce json
// @Param level body model.Level true "Level"
// @Success 201 {object} model.Level
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /level [post]
func (h *LevelHandler) CreateOrUpdateLevel(c echo.Context) error {
	level := new(model.Level)
	if err := c.Bind(level); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := h.levelUseCase.CreateOrUpdateLevel(level)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, level)
}

// DeleteLevel godoc
// @Summary Delete a level
// @Description Delete a level with the input payload
// @Tags levels
// @Accept json
// @Produce json
// @Param level body model.Level true "Level"
// @Success 200 {object} model.Level
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /level/delete [post]
func (h *LevelHandler) DeleteLevel(c echo.Context) error {
	level := new(model.Level)
	if err := c.Bind(level); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := h.levelUseCase.DeleteLevel(level)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, level)
}
