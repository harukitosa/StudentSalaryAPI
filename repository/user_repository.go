package repository

import "studentSalaryAPI/domain"

// UserRepository is interface
type UserRepository interface {
	Insert(user domain.User) (id int, err error)
	SelectByID(id int) (domain.User, error)
	SelectAll() ([]domain.User, error)
}
