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
		ID:           r.ID,
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
	res, err := r.db.NamedExec(`
	INSERT INTO reviews 
	(company_name,content,create_date_js,link,reasons,report,skill,user_name) 
	VALUES (:company_name,:content,:create_date_js,:link,:reasons,:report,:skill,:user_name)`,
		map[string]interface{}{
			"company_name":   review.CompanyName,
			"content":        review.Content,
			"create_date_js": review.CreateDateJS,
			"link":           review.Link,
			"reasons":        review.Reasons,
			"report":         review.Report,
			"skill":          review.Skill,
			"user_name":      review.UserName,
		})
	if err != nil {
		return 0, nil
	}
	var i int64
	i, err = res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	// あんまよくない
	return int(i), nil
}

func (r *reviewInfra) SelectByID(id int) (domain.Review, error) {
	query := `SELECT * FROM reviews WHERE id=:id`
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return domain.Review{}, nil
	}
	args := map[string]interface{}{
		"id": id,
	}
	var item review
	err = stmt.Get(&item, args)
	if err != nil {
		return domain.Review{}, nil
	}
	return item.toDomain(), nil
}

func (r *reviewInfra) SelectByName(name string) ([]domain.Review, error) {
	query := `SELECT * FROM reviews WHERE name=:name`
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, nil
	}
	args := map[string]interface{}{
		"name": name,
	}
	var items []review
	err = stmt.Select(&items, args)
	if err != nil {
		return nil, nil
	}

	var reviewsList []domain.Review
	for _, v := range items {
		reviewsList = append(reviewsList, v.toDomain())
	}
	return reviewsList, nil
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
