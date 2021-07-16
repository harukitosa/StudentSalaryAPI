package infra

import (
	"studentSalaryAPI/domain"

	"github.com/jmoiron/sqlx"
)

type workdataInfra struct {
	db *sqlx.DB
}

func NewWorkDataInfra(db *sqlx.DB) domain.WorkDataRepository {
	return &workdataInfra{db: db}
}

func (r *workdataInfra) Insert(review domain.WorkData) (id int, err error) {
	return 0, nil
}

func (r *workdataInfra) SelectByID(id int) (domain.WorkData, error) {
	return domain.WorkData{}, nil
}

func (r *workdataInfra) SelectByName(name string) ([]domain.WorkData, error) {
	return nil, nil
}

func (r *workdataInfra) SelectAll([]domain.WorkData, error) {
	return
}
