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

func (r *reviewHandler) GetReview(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return r.GetAllReview(c)
	}
	return r.SelectByName(c)
}

func (r *reviewHandler) GetAllReview(c echo.Context) error {
	reviews, err := r.repository.SelectAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	body := make([]reviewBody, len(reviews))
	for i, v := range reviews {
		body[i] = createBody(v)
	}
	return c.JSON(http.StatusOK, body)
}

func (r *reviewHandler) SelectByName(c echo.Context) error {
	name := c.QueryParam("name")
	reviews, err := r.repository.SelectByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	body := make([]reviewBody, len(reviews))
	for i, v := range reviews {
		body[i] = createBody(v)
	}
	return c.JSON(http.StatusOK, body)
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
	return c.JSON(http.StatusOK, createBody(review))
}

func (h *reviewHandler) CreateReview(c echo.Context) error {
	review := new(reviewPostBody)
	c.Bind(review)
	id, err := h.repository.Insert(domain.Review{
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

func (r *reviewHandler) GetReviewByCreated(c echo.Context) error {
	reviews, err := r.repository.GetNewReview()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	body := make([]reviewBody, len(reviews))
	for i, v := range reviews {
		body[i] = createBody(v)
	}
	return c.JSON(http.StatusOK, body)
}

type reviewPostBody struct {
	CompanyName  string `json:"company_name"`
	Content      string `json:"content"`
	CreateDateJS string `json:"create_date"`
	Link         string `json:"link"`
	Reasons      string `json:"reasons"`
	Report       string `json:"report"`
	Skill        string `json:"skill"`
	UserName     string `json:"user_name"`
}

type reviewBody struct {
	ID           int    `json:"ID"`
	CompanyName  string `json:"company_name"`
	Content      string `json:"content"`
	CreateDateJS string `json:"create_date"`
	Link         string `json:"link"`
	Reasons      string `json:"reasons"`
	Report       string `json:"report"`
	Skill        string `json:"skill"`
	UserName     string `json:"user_name"`
}

func createBody(review domain.Review) reviewBody {
	return reviewBody{
		ID:           review.ID,
		CompanyName:  review.CompanyName,
		Content:      review.Content,
		CreateDateJS: review.CreateDateJS,
		Link:         review.Link,
		Reasons:      review.Reasons,
		Report:       review.Report,
		Skill:        review.Skill,
		UserName:     review.UserName,
	}
}
