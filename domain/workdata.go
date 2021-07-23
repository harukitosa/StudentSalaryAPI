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

func NewWorkData(
	create_data_js *string,
	detail *string,
	experience *string,
	isShow *bool,
	name *string,
	salary *int,
	term *string,
	Type *string,
	workDays *string,
	workType *string) WorkData {
	return WorkData{
		Name:         convertNilString(name),
		Salary:       convertNilInt(salary),
		CreateDataJS: convertNilString(create_data_js),
		Detail:       convertNilString(detail),
		Experience:   convertNilString(experience),
		IsShow:       convertNilBoolean(isShow),
		Term:         convertNilString(term),
		Type:         convertNilString(Type),
		WorkDays:     convertNilString(workDays),
		WorkType:     convertNilString(workType),
	}
}

func convertNilString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func convertNilBoolean(b *bool) bool {
	if b == nil {
		return false
	}
	return true
}

func convertNilInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

type WorkDataRepository interface {
	Insert(salary WorkData) (id int, err error)
	SelectByID(id int) (WorkData, error)
	SelectByName(name string) ([]WorkData, error)
	SelectAll() ([]WorkData, error)
}
