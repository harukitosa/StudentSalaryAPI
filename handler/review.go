package handler

import (
	"net/http"
	"studentSalaryAPI/application"
	"studentSalaryAPI/handler/dto"
	"studentSalaryAPI/model"

	"github.com/labstack/echo/v4"
)

// ReviewHandler is
type ReviewHandler struct {
	reviewApplication application.ReviewApplication
}

// NewReviewHandler is
func NewReviewHandler(ReviewApplication application.ReviewApplication) ReviewHandler {
	return ReviewHandler{reviewApplication: ReviewApplication}
}

// CreateReview is
func (h *ReviewHandler) CreateReview(c echo.Context) error {
	review := new(dto.ReviewBody)
	c.Bind(review)
	id, err := h.reviewApplication.Insert(model.Review{
		CompanyName:  review.CompanyName,
		Content:      review.Content,
		Link:         review.Link,
		Reasons:      review.Reasons,
		Report:       review.Report,
		Skill:        review.Skill,
		CreateDateJS: review.CreateDateJS,
		UserName:     review.UserName,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, id)
}

// GetAllReview is
func (h *ReviewHandler) GetAllReview(c echo.Context) error {
	Review, err := h.reviewApplication.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Review)
}
