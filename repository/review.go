package repository

import "studentSalaryAPI/domain"

type ReviewRepository interface {
	Insert(review domain.Review) (id int, err error)
	SelectByID(id int) (domain.Review, error)
	SelectByName(name string) ([]domain.Review, error)
	SelectAll() ([]domain.Review, error)
	GetNewReview() ([]domain.Review, error)
}
