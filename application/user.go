package application

import (
	"studentSalaryAPI/model"
	"studentSalaryAPI/repository"
)

// UserApplication is
type UserApplication struct {
	useRepository repository.UserRepository
}

// NewUserApplication is
func NewUserApplication(repository repository.UserRepository) UserApplication {
	return UserApplication{
		useRepository: repository,
	}
}

// Insert is
func (u *UserApplication) Insert(user model.User) (int, error) {
	return u.useRepository.Insert(user)
}

// GetAll is
func (u *UserApplication) GetAll() ([]model.User, error) {
	return u.useRepository.SelectAll()
}
