package repository

import "studentSalaryAPI/model"

// ReviewRepository is interface
type ReviewRepository interface {
	Insert(user model.Review) (id int, err error)
	SelectAll() ([]model.Review, error)
}
