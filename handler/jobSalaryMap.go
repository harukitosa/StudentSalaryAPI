package handler

import (
	"net/http"
	"studentSalaryAPI/application"

	"github.com/labstack/echo/v4"
)

// JobSalaryMapHandler is
type JobSalaryMapHandler struct {
	jobSalaryMapApplication application.JobSalaryMapApplication
}

// NewJobSalaryMapHandler is
func NewJobSalaryMapHandler(jobSalaryMapApplication application.JobSalaryMapApplication) JobSalaryMapHandler {
	return JobSalaryMapHandler{jobSalaryMapApplication: jobSalaryMapApplication}
}

// GetJobSalaryMap is
func (h *JobSalaryMapHandler) GetJobSalaryMap(c echo.Context) error {
	jobSalaryMap, err := h.jobSalaryMapApplication.Select()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, jobSalaryMap)
}
