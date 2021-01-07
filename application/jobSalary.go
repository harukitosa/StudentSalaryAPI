package application

import (
	"studentSalaryAPI/model"
	"studentSalaryAPI/repository"
)

// JobSalaryApplication is
type JobSalaryApplication struct {
	jobSalaryepository repository.JobSalaryRepository
}

// NewJobSalaryApplication is
func NewJobSalaryApplication(repository repository.JobSalaryRepository) JobSalaryApplication {
	return JobSalaryApplication{
		jobSalaryepository: repository,
	}
}

// Insert is
func (u *JobSalaryApplication) Insert(jobSalary model.JobSalary) (int, error) {
	return u.jobSalaryepository.Insert(jobSalary)
}

// GetAll is
func (u *JobSalaryApplication) GetAll() ([]model.JobSalary, error) {
	return u.jobSalaryepository.SelectAll()
}
