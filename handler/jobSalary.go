package handler

import (
	"net/http"
	"studentSalaryAPI/application"
	"studentSalaryAPI/handler/dto"
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
		job := new(dto.JobSalaryBody)
		c.Bind(job)

		id, err := h.jobSalaryApplication.Insert(model.JobSalary{
			Name:         job.Name,
			CreateDataJS: job.CreateDataJS,
			Detail:       job.Detail,
			Experience:   job.Experience,
			IsShow:       false,
			Salary:       job.Salary,
			Term:         job.Term,
			Type:         job.Type,
			WorkDays:     job.WorkDays,
			WorkType:     job.WorkType,
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
