package repository

import "studentSalaryAPI/domain"

type BlogRepository interface {
	SelectByName(name string) (*domain.Blog, error)
	Select() ([]domain.Blog, error)
	GetCompanyNameList() ([]string, error)
}
