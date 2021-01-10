package application

import (
	"studentSalaryAPI/domain"
	"studentSalaryAPI/repository"
	"studentSalaryAPI/service"
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

// GetStatistics will return statistics of jobSalary
func (u *JobSalaryApplication) GetStatistics() (int, int, int, int, error) {
	list, err := u.jobSalaryepository.SelectAll()
	if err != nil {
		return 0, 0, 0, 0, err
	}
	jobSalaryService := new(service.JobSalaryService)
	count := len(list)
	avg := jobSalaryService.GetAvg(list)
	mid := jobSalaryService.GetMid(list)
	companyCount := jobSalaryService.GetCountByCompanyName(list)
	return count, avg, mid, companyCount, nil
}
