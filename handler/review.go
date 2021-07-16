package handler

import (
	"net/http"
	"strconv"
	"studentSalaryAPI/domain"

	"github.com/labstack/echo/v4"
)

type reviewHandler struct {
	repository domain.ReviewRepository
}

func NewReviewHandler(repository domain.ReviewRepository) *reviewHandler {
	return &reviewHandler{repository: repository}
}

func (r *reviewHandler) GetAllReview(c echo.Context) error {
	reviews, err := r.repository.SelectAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, reviews)
}

func (r *reviewHandler) GetReviewByID(c echo.Context) error {
	id := c.Param("id")
	formatID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	review, err := r.repository.SelectByID(formatID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, review)
}
