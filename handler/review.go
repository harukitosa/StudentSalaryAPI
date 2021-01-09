package handler

import (
	"log"
	"net/http"
	"strconv"
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
	name := c.QueryParam("name")
	if name != "" {
		response, err := h.reviewApplication.GetByName(name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if len(response) == 0 {
			return c.JSON(http.StatusNotFound, "not found")
		}
		return c.JSON(http.StatusOK, response)
	}

	review, err := h.reviewApplication.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, review)
}

// GetReviewByID is
func (h *ReviewHandler) GetReviewByID(c echo.Context) error {
	id := c.Param("id")
	formatID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	review, err := h.reviewApplication.GetByID(formatID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, review)
}

// ExportReview is
func (h *ReviewHandler) ExportReview(c echo.Context) error {
	reviews := &[]dto.ExportReviewBody{}
	err := c.Bind(reviews)
	if err != nil {
		log.Println(err)
	}
	for _, review := range *reviews {
		_, err := h.reviewApplication.Insert(model.Review{
			CompanyName:  review.CompanyName,
			Content:      review.Content,
			Link:         review.Link,
			Reasons:      review.Reasons,
			Report:       review.Report,
			Skill:        review.Skill,
			CreateDateJS: strconv.Itoa(review.CreateDateJS),
			UserName:     review.UserName,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, len(*reviews))
}

// GetReviewByCreated is
func (h *ReviewHandler) GetReviewByCreated(c echo.Context) error {
	jobSalaryMap, err := h.reviewApplication.GetByCreated()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, jobSalaryMap)
}
