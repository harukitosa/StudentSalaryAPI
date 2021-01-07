package handler

import (
	"net/http"
	"studentSalaryAPI/application"
	"studentSalaryAPI/model"

	"github.com/labstack/echo/v4"
)

// JobSalaryHandler is
type JobSalaryHandler struct {
	jobSalaryApplication application.JobSalaryApplication
}

// NewJobSalaryHandler is
func NewJobSalaryHandler(jobSalaryApplication application.JobSalaryApplication) JobSalaryHandler {
	return JobSalaryHandler{jobSalaryApplication: jobSalaryApplication}
}

// CreateJobSalary is
func (h *JobSalaryHandler) CreateJobSalary() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := h.jobSalaryApplication.Insert(model.JobSalary{
			Name:         "test",
			CreateDataJS: "12121313414",
			Detail:       "詳細",
			Experience:   "2~3年",
			IsShow:       true,
			Salary:       2000,
			Term:         "2weeks",
			Type:         "ios developer",
			WorkDays:     "2days",
			WorkType:     "インターン",
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, id)
	}
}

// GetAllJobSalary is
func (h *JobSalaryHandler) GetAllJobSalary() echo.HandlerFunc {
	return func(c echo.Context) error {
		JobSalary, err := h.jobSalaryApplication.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, JobSalary)
	}
}
