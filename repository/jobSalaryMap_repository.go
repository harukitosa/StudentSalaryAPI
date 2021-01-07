package repository

import "studentSalaryAPI/model"

// JobSalaryMapRepository is interface
type JobSalaryMapRepository interface {
	Select() ([]model.JobSalaryMap, error)
}
