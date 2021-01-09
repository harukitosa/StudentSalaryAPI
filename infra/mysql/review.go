package mysql

import (
	"studentSalaryAPI/model"
	"studentSalaryAPI/repository"

	"gorm.io/gorm"
)

// ReviewRepositoryImpl is dependent sqlx and UserRepository
type ReviewRepositoryImpl struct {
	db *gorm.DB
}

// NewReviewRepositoryImpl is
func NewReviewRepositoryImpl(DB *gorm.DB) repository.ReviewRepository {
	return &ReviewRepositoryImpl{
		db: DB,
	}
}

// Insert is
func (r *ReviewRepositoryImpl) Insert(review model.Review) (int, error) {
	result := r.db.Create(&review)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(review.ID), nil
}

// SelectAll is
func (r *ReviewRepositoryImpl) SelectAll() ([]model.Review, error) {
	var reviews []model.Review
	result := r.db.Find(&reviews)
	return reviews, result.Error
}

// SelectByID is
func (r *ReviewRepositoryImpl) SelectByID(id int) (model.Review, error) {
	var review model.Review
	result := r.db.First(&review, id)
	return review, result.Error
}

// SelectByName is
func (r *ReviewRepositoryImpl) SelectByName(name string) ([]model.Review, error) {
	var reviews []model.Review
	result := r.db.Where("company_name == ?", name).Find(&reviews)
	return reviews, result.Error
}
