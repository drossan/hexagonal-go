package api

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

// LevelPrivilegesHandler manages level privileges
type LevelPrivilegesHandler struct {
	levelPrivilegesUseCase *usecase.LevelPrivilegesUseCase
}

// NewLevelPrivilegesHandler initializes a new LevelPrivilegesHandler
func NewLevelPrivilegesHandler(e *echo.Echo, uc *usecase.LevelPrivilegesUseCase) *LevelPrivilegesHandler {
	return &LevelPrivilegesHandler{levelPrivilegesUseCase: uc}
}

// RegisterRoutes registers level privileges routes
func (h *LevelPrivilegesHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/level-privileges", h.GetAllLevelPrivileges)
	g.POST("/level-privilege", h.CreateLevelPrivilege)
	g.POST("/level-privilege/delete", h.DeleteLevelPrivilege)
}

// GetAllLevelPrivileges godoc
// @Summary Get all level privileges
// @Description Get all level privileges
// @Tags level-privileges
// @Accept json
// @Produce json
// @Success 200 {object} []model.LevelPrivileges
// @Failure 500 {object} map[string]interface{}
// @Router /level-privileges [get]
func (h *LevelPrivilegesHandler) GetAllLevelPrivileges(c echo.Context) error {
	levelPrivileges, err := h.levelPrivilegesUseCase.GetAllLevelPrivilege()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, levelPrivileges)
}

// CreateLevelPrivilege godoc
// @Summary Create a new level privilege
// @Description Create a new level privilege with the input payload
// @Tags level-privileges
// @Accept json
// @Produce json
// @Param levelPrivilege body model.LevelPrivileges true "Level Privilege"
// @Success 201 {object} model.LevelPrivileges
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /level-privilege [post]
func (h *LevelPrivilegesHandler) CreateLevelPrivilege(c echo.Context) error {
	levelPrivilege := new(model.LevelPrivileges)
	if err := c.Bind(levelPrivilege); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := h.levelPrivilegesUseCase.CreateOrUpdateLevelPrivilege(levelPrivilege)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, levelPrivilege)
}

// DeleteLevelPrivilege godoc
// @Summary Delete a level privilege
// @Description Delete a level privilege with the input payload
// @Tags level-privileges
// @Accept json
// @Produce json
// @Param levelPrivilege body model.LevelPrivileges true "Level Privilege"
// @Success 200 {object} model.LevelPrivileges
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /level-privilege/delete [post]
func (h *LevelPrivilegesHandler) DeleteLevelPrivilege(c echo.Context) error {
	levelPrivilege := new(model.LevelPrivileges)
	if err := c.Bind(levelPrivilege); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := h.levelPrivilegesUseCase.DeleteLevelPrivilege(levelPrivilege)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, levelPrivilege)
}
