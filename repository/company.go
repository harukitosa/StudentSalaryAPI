package repository

import "studentSalaryAPI/domain"

type CompanyRepository interface {
	SelectByTop() ([]domain.Company, error)
	SelectByName(name string) (*domain.Company, error)
	Select() ([]domain.Company, error)
}
