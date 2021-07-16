package infra

import (
	"studentSalaryAPI/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type reviewInfra struct {
	db *sqlx.DB
}

func NewReviewInfra(db *sqlx.DB) domain.ReviewRepository {
	return &reviewInfra{db: db}
}

type review struct {
	ID           int        `db:"id"`
	CreatedAt    *time.Time `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
	CompanyName  string     `db:"company_name"`
	Content      string     `db:"content"`
	CreateDateJS string     `db:"create_date_js"`
	Link         string     `db:"link"`
	Reasons      string     `db:"reasons"`
	Report       string     `db:"report"`
	Skill        string     `db:"skill"`
	UserName     string     `db:"user_name"`
}

func (r *review) toDomain() domain.Review {
	return domain.Review{
		CompanyName:  r.CompanyName,
		Content:      r.Content,
		CreateDateJS: r.CreateDateJS,
		Link:         r.Link,
		Reasons:      r.Reasons,
		Report:       r.Report,
		Skill:        r.Skill,
		UserName:     r.UserName,
	}
}

func (r *reviewInfra) Insert(review domain.Review) (id int, err error) {
	return 0, nil
}

func (r *reviewInfra) SelectByID(id int) (domain.Review, error) {
	query := `SELECT * FROM reviews WHERE id=?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return domain.Review{}, nil
	}
	row := stmt.QueryRow(id)
	review := new(review)
	err = row.Scan(&review)
	if err != nil {
		return domain.Review{}, err
	}
	return review.toDomain(), nil
}

func (r *reviewInfra) SelectByName(name string) ([]domain.Review, error) {
	return nil, nil
}

func (r *reviewInfra) SelectAll() ([]domain.Review, error) {
	rows, err := r.db.Queryx("SELECT * FROM reviews")
	if err != nil {
		return nil, err
	}

	var reviews []review
	for rows.Next() {
		review := new(review)
		err := rows.StructScan(&review)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, *review)
	}

	var reviewsList []domain.Review
	for _, v := range reviews {
		reviewsList = append(reviewsList, v.toDomain())
	}
	return reviewsList, nil
}
