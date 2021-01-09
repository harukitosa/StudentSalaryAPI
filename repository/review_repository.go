package repository

import "studentSalaryAPI/domain"

// ReviewRepository is interface
type ReviewRepository interface {
	Insert(user domain.Review) (id int, err error)
	SelectAll() ([]domain.Review, error)
	SelectByID(id int) (domain.Review, error)
	SelectByName(name string) ([]domain.Review, error)
	SelectByCreated() ([]domain.Review, error)
}
