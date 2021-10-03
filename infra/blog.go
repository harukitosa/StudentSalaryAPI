package infra

import (
	"fmt"
	"studentSalaryAPI/domain"
	"studentSalaryAPI/repository"

	"github.com/jmoiron/sqlx"
)

type blogInfra struct {
	db *sqlx.DB
}

func NewBlogInfra(db *sqlx.DB) repository.BlogRepository {
	return &blogInfra{db: db}
}

func (b *blogInfra) SelectByName(name string) (*domain.Blog, error) {
	return nil, fmt.Errorf("not impl")
}

func (b *blogInfra) Select() ([]domain.Blog, error) {
	query := `select * from blogs order by year DESC`
	rows, err := b.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	var list []blog
	for rows.Next() {
		c := new(blog)
		err := rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		list = append(list, *c)
	}

	var res []domain.Blog
	for _, v := range list {
		res = append(res, domain.Blog{
			ID:           v.ID,
			URL:          v.URL,
			Title:        v.Title,
			Company_name: v.Company_name,
			Year:         v.Year,
			Season:       v.Season})
	}
	return res, nil
}

type NameList struct {
	Name string `db:"company_name"`
}

func (b *blogInfra) GetCompanyNameList() ([]string, error) {
	query := `SELECT company_name FROM blogs GROUP BY company_name ORDER BY company_name ASC`

	rows, err := b.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	var list []NameList
	for rows.Next() {
		data := new(NameList)
		err := rows.StructScan(&data)
		if err != nil {
			return nil, err
		}
		list = append(list, *data)
	}
	var res []string
	for _, v := range list {
		res = append(res, v.Name)
	}
	return res, nil
}

type blog struct {
	ID           int    `db:"id"`
	URL          string `db:"url"`
	Title        string `db:"title"`
	Company_name string `db:"company_name"`
	Year         int    `db:"year"`
	Season       string `db:"season"`
}
