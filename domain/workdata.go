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

// func NewWorkData(id *int,
// 	create_data_js *string,
// 	detail *string,
// 	experience *string,
// 	isShow *bool,
// 	name *string,
// 	salary *int,
// 	term *string,
// 	Type *string,
// 	workDays *string,
// 	workType *string) WorkData {
// 	return WorkData{
// 		Name:         *name,
// 		Salary:       *salary,
// 		CreateDataJS: *create_data_js,
// 		Detail:       *detail,
// 		Experience:   *experience,
// 		IsShow:       *isShow,
// 		Term:         *term,
// 		Type:         *Type,
// 		WorkDays:     *workDays,
// 		WorkType:     *workType,
// 	}
// }

type WorkDataRepository interface {
	Insert(salary WorkData) (id int, err error)
	SelectByID(id int) (WorkData, error)
	SelectByName(name string) ([]WorkData, error)
	SelectAll() ([]WorkData, error)
}
