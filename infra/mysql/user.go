package mysql

import (
	"studentSalaryAPI/domain"
	"studentSalaryAPI/repository"

	"gorm.io/gorm"
)

// UserRepositoryImpl is dependent sqlx and UserRepository
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepositoryImpl is
func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		db: DB,
	}
}

// Insert is
func (r *UserRepositoryImpl) Insert(user domain.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.ID), nil
}

// SelectByID is
func (r *UserRepositoryImpl) SelectByID(id int) (domain.User, error) {
	var user domain.User
	tx := r.db.First(&user, id)
	return user, tx.Error
}

// SelectAll is
func (r *UserRepositoryImpl) SelectAll() ([]domain.User, error) {
	var users []domain.User
	result := r.db.Find(&users)
	return users, result.Error
}
