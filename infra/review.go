package infra

import (
	"log"
	"studentSalaryAPI/domain"
	"studentSalaryAPI/repository"
	"time"

	"github.com/jmoiron/sqlx"
)

type reviewInfra struct {
	db *sqlx.DB
}

func NewReviewInfra(db *sqlx.DB) repository.ReviewRepository {
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

func (r *review) toDomain() (domain.Review, error) {
	return domain.NewReview(
		&r.ID,
		&r.CompanyName,
		&r.Content,
		&r.CreateDateJS,
		&r.Link,
		&r.Reasons,
		&r.Report,
		&r.Skill,
		&r.UserName,
	)
}

func (r *reviewInfra) Insert(review domain.Review) (id int, err error) {
	res, err := r.db.NamedExec(`
	INSERT INTO reviews 
	(company_name,content,create_date_js,link,reasons,report,skill,user_name) 
	VALUES (:company_name,:content,:create_date_js,:link,:reasons,:report,:skill,:user_name)`,
		map[string]interface{}{
			"company_name":   review.GetCompanyName().String(),
			"content":        review.GetContent().String(),
			"create_date_js": review.GetCreateDate().String(),
			"link":           review.GetLink().String(),
			"reasons":        review.GetReasons().String(),
			"report":         review.GetReport().String(),
			"skill":          review.GetSkill().String(),
			"user_name":      review.GetUserName().String(),
		})
	if err != nil {
		return 0, err
	}
	var i int64
	i, err = res.LastInsertId()
	if err != nil {
		return 0, err
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
	return item.toDomain()
}

func (r *reviewInfra) SelectByName(name string) ([]domain.Review, error) {
	query := `SELECT * FROM reviews WHERE company_name=:name`
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
		return nil, err
	}
	var reviewsList []domain.Review
	for _, v := range items {
		i, err := v.toDomain()
		if err != nil {
			return nil, err
		}
		reviewsList = append(reviewsList, i)
	}
	return reviewsList, nil
}

func (r *reviewInfra) SelectAll() ([]domain.Review, error) {
	var rows *sqlx.Rows
	var err error
	rows, err = r.db.Queryx("SELECT * FROM reviews")
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
		i, err := v.toDomain()
		if err != nil {
			return nil, err
		}
		reviewsList = append(reviewsList, i)
	}
	return reviewsList, nil
}

func (r *reviewInfra) GetNewReview() ([]domain.Review, error) {
	rows, err := r.db.Queryx(`
		select * from reviews order by created_at DESC limit 3;
	`)
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
	log.Println(reviews)

	var reviewsList []domain.Review
	for _, v := range reviews {
		i, err := v.toDomain()
		if err != nil {
			return nil, err
		}
		reviewsList = append(reviewsList, i)
	}
	return reviewsList, nil
}
