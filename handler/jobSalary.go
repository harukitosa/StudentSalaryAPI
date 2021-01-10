package handler

import (
	"log"
	"net/http"
	"strconv"
	"studentSalaryAPI/application"
	"studentSalaryAPI/domain"
	"studentSalaryAPI/handler/dto"

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
func (h *JobSalaryHandler) CreateJobSalary(c echo.Context) error {
	job := new(dto.JobSalaryBody)
	c.Bind(job)

	id, err := h.jobSalaryApplication.Insert(domain.JobSalary{
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

// GetAllJobSalary is
func (h *JobSalaryHandler) GetAllJobSalary(c echo.Context) error {
	name := c.QueryParam("name")
	if name != "" {
		response, err := h.jobSalaryApplication.GetByName(name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if len(response) == 0 {
			return c.JSON(http.StatusNotFound, "not found")
		}
		return c.JSON(http.StatusOK, response)
	}

	JobSalary, err := h.jobSalaryApplication.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, JobSalary)
}

// GetStatistics is
func (h *JobSalaryHandler) GetStatistics(c echo.Context) error {
	count, avg, mid, companyCount, err := h.jobSalaryApplication.GetStatistics()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	s := dto.JobSalaryStatistics{Count: count, CompanyCount: companyCount, Avg: avg, Mid: mid}
	return c.JSON(http.StatusOK, s)
}

// ExportJobsSalary is
func (h *JobSalaryHandler) ExportJobsSalary(c echo.Context) error {
	jobs := &[]dto.ExportJobSalaryBody{}
	err := c.Bind(jobs)
	if err != nil {
		log.Fatal(err)
	}

	for _, job := range *jobs {
		salary, err := strconv.Atoi(job.Salary)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		_, err = h.jobSalaryApplication.Insert(domain.JobSalary{
			Name:         job.Name,
			CreateDataJS: strconv.Itoa(job.CreateDataJS),
			Detail:       job.Detail,
			Experience:   job.Experience,
			IsShow:       true,
			Salary:       salary,
			Term:         job.Term,
			Type:         job.Type,
			WorkDays:     job.WorkDays,
			WorkType:     job.WorkType,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, len(*jobs))
}
