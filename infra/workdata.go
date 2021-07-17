package infra

import (
	"studentSalaryAPI/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type workdataInfra struct {
	db *sqlx.DB
}

func NewWorkDataInfra(db *sqlx.DB) domain.WorkDataRepository {
	return &workdataInfra{db: db}
}

type workdata struct {
	ID           int        `db:"id"`
	CreatedAt    *time.Time `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
	CreateDataJS string     `db:"create_data_js"`
	Detail       string     `db:"detail"`
	Experience   string     `db:"experience"`
	IsShow       bool       `db:"is_show"`
	Name         string     `db:"name"`
	Salary       int        `db:"salary"`
	Term         string     `db:"term"`
	Type         string     `db:"type"`
	WorkDays     string     `db:"work_days"`
	WorkType     string     `db:"work_type"`
}

func (w *workdata) toDomain() domain.WorkData {
	return domain.WorkData{
		ID:           w.ID,
		Detail:       w.Detail,
		CreateDataJS: w.CreateDataJS,
		Experience:   w.Experience,
		IsShow:       w.IsShow,
		Name:         w.Name,
		Salary:       w.Salary,
		Term:         w.Term,
		Type:         w.Type,
		WorkDays:     w.WorkDays,
		WorkType:     w.WorkType,
	}
}

func (r *workdataInfra) Insert(review domain.WorkData) (id int, err error) {
	res, err := r.db.NamedExec(`
	INSERT INTO job_salaries
	(created_at,updated_at,create_data_js,detail,experience,is_show,name,salary,term,type,work_days,work_type)
	VALUES (:created_at,:updated_at,:create_data_js,:detail,:experience,:is_show,:name,:salary,:term,:type,:work_days,:work_type)`,
		map[string]interface{}{
			"created_at":     time.Now(),
			"updated_at":     time.Now(),
			"create_data_js": review.CreateDataJS,
			"detail":         review.Detail,
			"experience":     review.Experience,
			"is_show":        review.IsShow,
			"name":           review.Name,
			"salary":         review.Salary,
			"term":           review.Term,
			"type":           review.Type,
			"work_days":      review.WorkDays,
			"work_type":      review.WorkType,
		})
	if err != nil {
		return 0, err
	}
	i, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func (r *workdataInfra) SelectByID(id int) (domain.WorkData, error) {
	return domain.WorkData{}, nil
}

func (r *workdataInfra) SelectByName(name string) ([]domain.WorkData, error) {
	query := `SELECT * FROM job_salaries WHERE name=:name`
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, nil
	}
	args := map[string]interface{}{
		"name": name,
	}
	var items []workdata
	err = stmt.Select(&items, args)
	if err != nil {
		return nil, err
	}
	var workdataList []domain.WorkData
	for _, v := range items {
		workdataList = append(workdataList, v.toDomain())
	}
	return workdataList, nil
}

func (r *workdataInfra) SelectAll() ([]domain.WorkData, error) {
	rows, err := r.db.Queryx("SELECT * FROM job_salaries ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}

	var list []workdata
	for rows.Next() {
		data := new(workdata)
		err := rows.StructScan(&data)
		if err != nil {
			return nil, err
		}
		list = append(list, *data)
	}

	var workdataList []domain.WorkData
	for _, v := range list {
		workdataList = append(workdataList, v.toDomain())
	}
	return workdataList, nil
}
