package handler

import (
	"net/http"
	"studentSalaryAPI/application"
	"studentSalaryAPI/model"

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
func (h *UserHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := h.userApplication.Insert(model.User{Name: "test", ID: 1})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, id)
	}
}

// GetAllUser is
func (h *UserHandler) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := h.userApplication.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, user)
	}
}
