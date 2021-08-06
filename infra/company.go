package infra

import (
	"studentSalaryAPI/domain"

	"github.com/jmoiron/sqlx"
)

type companyInfra struct {
	db *sqlx.DB
}

type company struct {
	Name  string `db:"name"`
	Max   int    `db:"max"`
	Min   int    `db:"min"`
	Count int    `db:"count"`
}

func NewCompanyInfra(db *sqlx.DB) domain.CompanyRepository {
	return &companyInfra{db: db}
}
func (r *companyInfra) SelectByTop() ([]domain.Company, error) {
	query := `select name, max(salary) as max, min(salary) as min, count(*) as count from job_salaries group by name order by count DESC limit 3`
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	var list []company
	for rows.Next() {
		c := new(company)
		err := rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		list = append(list, *c)
	}

	var res []domain.Company
	for _, v := range list {
		res = append(res, domain.Company{Max: v.Max, Count: v.Count, Min: v.Min, Name: v.Name})
	}
	return res, nil
}
