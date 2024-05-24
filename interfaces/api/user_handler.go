package api

import (
	"github.com/drossan/core-api/helpers"
	"net/http"
	"strconv"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(e *echo.Echo, uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: uc}
}

func (h *UserHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/users/:page", h.PaginateUsers)
	g.GET("/user", h.GetUserData)
	g.POST("/user", h.CreateOrUpdateUser)
	g.POST("/user/delete", h.DeleteUser)
}

func (h *UserHandler) AuthRoutes(e *echo.Group) {
	e.POST("/login", h.Login)
}

// CreateOrUpdateUser godoc
// @Summary Create or update a user
// @Description Create or update a user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user [post]
func (h *UserHandler) CreateOrUpdateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if user.ID != 0 {
		err := h.userUseCase.UpdateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
	} else {
		err := h.userUseCase.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusCreated, user)
}

// GetUserData obtiene los datos del usuario actual
// @Summary Get current user data
// @Description Get data for the current user based on the token
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user [get]
func (h *UserHandler) GetUserData(c echo.Context) error {
	userID := helpers.GetCurrentUser(c)

	user, err := h.userUseCase.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

// PaginateUsers godoc
// @Summary Get users with pagination
// @Description Get users with pagination
// @Tags users
// @Accept json
// @Produce json
// @Param page path int true "Page number"
// @Param rows query int false "Rows per page"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{page} [get]
func (h *UserHandler) PaginateUsers(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil || page < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid page number"})
	}

	rows, err := strconv.Atoi(c.QueryParam("rows"))
	if err != nil || rows < 1 {
		rows = 50
	}

	users, total, err := h.userUseCase.PaginateUsers(page, rows)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": users,
		"total": total,
	})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/delete [post]
func (h *UserHandler) DeleteUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	err := h.userUseCase.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

// Login godoc
// @Summary Login a user
// @Description Login a user with the input payload
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func (h *UserHandler) Login(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	token, err := h.userUseCase.Login(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": 200,
		"data": echo.Map{
			"token": token,
		},
	})
}
