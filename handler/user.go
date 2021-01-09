package handler

import (
	"net/http"
	"studentSalaryAPI/application"
	"studentSalaryAPI/domain"

	"github.com/labstack/echo/v4"
)

// UserHandler is
type UserHandler struct {
	userApplication application.UserApplication
}

// NewUserHandler is
func NewUserHandler(userApplication application.UserApplication) UserHandler {
	return UserHandler{userApplication: userApplication}
}

// CreateUser is
func (h *UserHandler) CreateUser(c echo.Context) error {
	id, err := h.userApplication.Insert(domain.User{Name: "test"})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, id)
}

// GetAllUser is
func (h *UserHandler) GetAllUser(c echo.Context) error {
	user, err := h.userApplication.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}
