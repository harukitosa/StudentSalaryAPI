package repository

import "studentSalaryAPI/domain"

// JobSalaryRepository is interface
type JobSalaryRepository interface {
	Insert(user domain.JobSalary) (id int, err error)
	SelectByID(id int) (domain.JobSalary, error)
	SelectByName(name string) ([]domain.JobSalary, error)
	SelectAll() ([]domain.JobSalary, error)
}
