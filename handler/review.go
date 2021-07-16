package handler

import (
	"net/http"
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
