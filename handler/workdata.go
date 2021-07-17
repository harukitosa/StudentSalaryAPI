package handler

import (
	"net/http"
	"studentSalaryAPI/domain"

	"github.com/labstack/echo/v4"
)

type workdataHandler struct {
	repository domain.WorkDataRepository
}

func NewWorkDataHandler(repository domain.WorkDataRepository) *workdataHandler {
	return &workdataHandler{repository: repository}
}

func (r *workdataHandler) GetReview(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return r.GetAllReview(c)
	}
	return r.GetReviewByName(c)
}

func (r *workdataHandler) GetAllReview(c echo.Context) error {
	data, err := r.repository.SelectAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	body := make([]workDataBody, len(data))
	for i, v := range data {
		body[i] = createWorkDataBody(v)
	}
	return c.JSON(http.StatusOK, body)
}

func (r *workdataHandler) GetReviewByName(c echo.Context) error {
	name := c.QueryParam("name")
	data, err := r.repository.SelectByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	body := make([]workDataBody, len(data))
	for i, v := range data {
		body[i] = createWorkDataBody(v)
	}
	return c.JSON(http.StatusOK, body)
}

type workDataBody struct {
	ID           int    `json:"id"`
	CreateDataJS string `json:"create_data_js"`
	Detail       string `json:"detail"`
	Experience   string `json:"experience"`
	IsShow       bool   `json:"is_show"`
	Name         string `json:"name"`
	Salary       int    `json:"salary"`
	Term         string `json:"term"`
	Type         string `json:"type"`
	WorkDays     string `json:"work_days"`
	WorkType     string `json:"work_type"`
}

func createWorkDataBody(workdata domain.WorkData) workDataBody {
	return workDataBody{
		ID:           workdata.ID,
		CreateDataJS: workdata.CreateDataJS,
		Detail:       workdata.Detail,
		Experience:   workdata.Experience,
		IsShow:       workdata.IsShow,
		Name:         workdata.Name,
		Salary:       workdata.Salary,
		Term:         workdata.Term,
		Type:         workdata.Type,
		WorkDays:     workdata.WorkDays,
		WorkType:     workdata.WorkType,
	}
}
