package infra

import (
	"studentSalaryAPI/domain"

	"github.com/jmoiron/sqlx"
)

type reviewInfra struct {
	db *sqlx.DB
}

func NewReviewInfra(db *sqlx.DB) domain.ReviewRepository {
	return &reviewInfra{db: db}
}

func (r *reviewInfra) Insert(review domain.Review) (id int, err error) {
	return 0, nil
}

func (r *reviewInfra) SelectByID(id int) (domain.Review, error) {
	return domain.Review{}, nil
}

func (r *reviewInfra) SelectByName(name string) ([]domain.Review, error) {
	return nil, nil
}

func (r *reviewInfra) SelectAll([]domain.Review, error) {
	return
}
