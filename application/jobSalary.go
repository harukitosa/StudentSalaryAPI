package application

import (
	"studentSalaryAPI/domain"
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
func (u *JobSalaryApplication) Insert(jobSalary domain.JobSalary) (int, error) {
	return u.jobSalaryepository.Insert(jobSalary)
}

// GetAll is
func (u *JobSalaryApplication) GetAll() ([]domain.JobSalary, error) {
	return u.jobSalaryepository.SelectAll()
}

// GetByName is
func (u *JobSalaryApplication) GetByName(name string) ([]domain.JobSalary, error) {
	return u.jobSalaryepository.SelectByName(name)
}
