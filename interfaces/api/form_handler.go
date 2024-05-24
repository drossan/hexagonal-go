package api

import (
	"net/http"
	"strconv"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/usecase"
	"github.com/labstack/echo/v4"
)

type FormHandler struct {
	formUseCase *usecase.FormUseCase
}

func NewFormHandler(e *echo.Echo, uc *usecase.FormUseCase) *FormHandler {
	return &FormHandler{formUseCase: uc}
}

func (h *FormHandler) RegisterRoutes(g *echo.Group) {
	g.POST("/form", h.CreateOrUpdateForm)
	g.GET("/forms", h.GetAllForms)
	g.GET("/forms/:page", h.PaginateForms)
	g.POST("/form/delete", h.DeleteForm)
}

// CreateOrUpdateForm godoc
// @Summary Create or update a form
// @Description Create or update a form with the input payload
// @Tags forms
// @Accept json
// @Produce json
// @Param form body model.Form true "Form"
// @Success 201 {object} model.Form
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /form [post]
func (h *FormHandler) CreateOrUpdateForm(c echo.Context) error {
	form := new(model.Form)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := h.formUseCase.CreateOrUpdateForm(form)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, form)
}

// GetAllForms godoc
// @Summary Get all forms
// @Description Get all forms
// @Tags forms
// @Accept json
// @Produce json
// @Success 200 {array} model.Form
// @Failure 500 {object} map[string]interface{}
// @Router /forms [get]
func (h *FormHandler) GetAllForms(c echo.Context) error {
	forms, err := h.formUseCase.GetAllForms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, forms)
}

// PaginateForms godoc
// @Summary Get forms with pagination
// @Description Get forms with pagination
// @Tags forms
// @Accept json
// @Produce json
// @Param page path int true "Page number"
// @Param rows query int false "Rows per page"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /forms/{page} [get]
func (h *FormHandler) PaginateForms(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil || page < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid page number"})
	}

	rows, err := strconv.Atoi(c.QueryParam("rows"))
	if err != nil || rows < 1 {
		rows = 50
	}

	forms, total, err := h.formUseCase.PaginateForms(page, rows)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": forms,
		"total": total,
	})
}

// DeleteForm godoc
// @Summary Delete a form
// @Description Delete a form with the input payload
// @Tags forms
// @Accept json
// @Produce json
// @Param form body model.Form true "Form"
// @Success 200 {object} model.Form
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /form/delete [post]
func (h *FormHandler) DeleteForm(c echo.Context) error {
	form := new(model.Form)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := h.formUseCase.DeleteForm(form)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, form)
}
