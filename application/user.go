package application

import (
	"studentSalaryAPI/domain"
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
func (u *UserApplication) Insert(user domain.User) (int, error) {
	return u.useRepository.Insert(user)
}

// GetAll is
func (u *UserApplication) GetAll() ([]domain.User, error) {
	return u.useRepository.SelectAll()
}
