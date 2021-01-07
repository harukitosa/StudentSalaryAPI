package mysql

import (
	"studentSalaryAPI/model"
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
func (r *UserRepositoryImpl) Insert(user model.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.ID), nil
}

// SelectByID is
func (r *UserRepositoryImpl) SelectByID(id int) (model.User, error) {
	var user model.User
	tx := r.db.First(&user, id)
	return user, tx.Error
}

// SelectAll is
func (r *UserRepositoryImpl) SelectAll() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	return users, result.Error
}
