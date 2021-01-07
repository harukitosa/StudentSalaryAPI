package repository

import "studentSalaryAPI/model"

// JobSalaryRepository is interface
type JobSalaryRepository interface {
	Insert(user model.JobSalary) (id int, err error)
	SelectByID(id int) (model.JobSalary, error)
	SelectAll() ([]model.JobSalary, error)
}
