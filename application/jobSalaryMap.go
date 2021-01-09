package application

import (
	"studentSalaryAPI/domain"
	"studentSalaryAPI/repository"
)

// JobSalaryMapApplication is
type JobSalaryMapApplication struct {
	jobSalaryMaprepository repository.JobSalaryMapRepository
}

// NewJobSalaryMapApplication is
func NewJobSalaryMapApplication(repository repository.JobSalaryMapRepository) JobSalaryMapApplication {
	return JobSalaryMapApplication{
		jobSalaryMaprepository: repository,
	}
}

// Select is
func (u *JobSalaryMapApplication) Select() ([]domain.JobSalaryMap, error) {
	return u.jobSalaryMaprepository.Select()
}

// SelectByCount is
func (u *JobSalaryMapApplication) SelectByCount() ([]domain.JobSalaryMap, error) {
	return u.jobSalaryMaprepository.SelectByCount()
}
