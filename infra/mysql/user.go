package mysql

import (
	"errors"
	"studentSalaryAPI/model"
	"studentSalaryAPI/repository"

	"github.com/jmoiron/sqlx"
)

// UserRepositoryImpl is dependent sqlx and UserRepository
type UserRepositoryImpl struct {
	db *sqlx.DB
}

// NewUserRepositoryImpl is
func NewUserRepositoryImpl(DB *sqlx.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		db: DB,
	}
}

// Insert is
func (r *UserRepositoryImpl) Insert(user model.User) (int, error) {
	result, err := r.db.NamedExec("INSERT INTO user (name) VALUES (:name)", user)
	if err != nil {
		return 0, err
	}
	var id int64
	id, err = result.LastInsertId()
	return int(id), err
}

// SelectByID is
func (r *UserRepositoryImpl) SelectByID(id int) (model.User, error) {
	return model.User{ID: 0, Name: "hoge"}, errors.New("Not Impl")
}

// SelectAll is
func (r *UserRepositoryImpl) SelectAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, "SELECT * FROM user ORDER BY id ASC")
	return users, err
}
