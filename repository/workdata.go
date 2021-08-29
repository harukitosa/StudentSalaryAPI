package repository

import "studentSalaryAPI/domain"

type WorkDataRepository interface {
	Insert(domain.WorkData) (id int, err error)
	SelectByID(id int) (domain.WorkData, error)
	SelectByName(name string) ([]domain.WorkData, error)
	SelectAll() ([]domain.WorkData, error)
}
