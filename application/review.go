package application

import (
	"studentSalaryAPI/domain"
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
func (a *ReviewApplication) Insert(review domain.Review) (int, error) {
	return a.reviewRepository.Insert(review)
}

// GetAll is
func (a *ReviewApplication) GetAll() ([]domain.Review, error) {
	return a.reviewRepository.SelectAll()
}

// GetByID is
func (a *ReviewApplication) GetByID(id int) (domain.Review, error) {
	return a.reviewRepository.SelectByID(id)
}

// GetByName is
func (a *ReviewApplication) GetByName(name string) ([]domain.Review, error) {
	return a.reviewRepository.SelectByName(name)
}

// GetByCreated is
func (a *ReviewApplication) GetByCreated() ([]domain.Review, error) {
	return a.reviewRepository.SelectByCreated()
}
