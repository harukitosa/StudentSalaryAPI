package repository

import "studentSalaryAPI/domain"

// JobSalaryMapRepository is interface
type JobSalaryMapRepository interface {
	Select() ([]domain.JobSalaryMap, error)
	SelectByCount() ([]domain.JobSalaryMap, error)
}
