package domain

type WorkData struct {
	ID           int
	CreateDataJS string
	Detail       string
	Experience   string
	IsShow       bool
	Name         string
	Salary       int
	Term         string
	Type         string
	WorkDays     string
	WorkType     string
}

type WorkDataRepository interface {
	Insert(salary WorkData) (id int, err error)
	SelectByID(id int) (WorkData, error)
	SelectByName(name string) ([]WorkData, error)
	SelectAll() ([]WorkData, error)
}
