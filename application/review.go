package application

import (
	"studentSalaryAPI/model"
	"studentSalaryAPI/repository"
)

// ReviewApplication is
type ReviewApplication struct {
	reviewRepository repository.ReviewRepository
}

// NewReviewApplication is
func NewReviewApplication(repository repository.ReviewRepository) ReviewApplication {
	return ReviewApplication{
		reviewRepository: repository,
	}
}

// Insert is
func (a *ReviewApplication) Insert(review model.Review) (int, error) {
	return a.reviewRepository.Insert(review)
}

// GetAll is
func (a *ReviewApplication) GetAll() ([]model.Review, error) {
	return a.reviewRepository.SelectAll()
}

// GetByID is
func (a *ReviewApplication) GetByID(id int) (model.Review, error) {
	return a.reviewRepository.SelectByID(id)
}
