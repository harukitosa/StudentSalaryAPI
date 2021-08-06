package domain

import "sort"

// domain
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

type WorkDataService struct{}

func (s *WorkDataService) GetSalaryAvg(list []WorkData) int {
	if len(list) == 0 {
		return 0
	}
	ans := 0
	for _, v := range list {
		ans += v.Salary
	}
	return ans / len(list)
}

func (s *WorkDataService) GetSalaryMid(list []WorkData) int {
	if len(list) == 0 {
		return 0
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Salary < list[j].Salary })
	i := len(list) / 2
	if len(list)%2 == 0 {
		return (list[i].Salary + list[i-1].Salary) / 2
	}
	return list[i].Salary
}

// GetCountByCompanyName ユニークな会社名がいくつあるのか調べる
func (s *WorkDataService) GetCountByCompanyName(list []WorkData) int {
	if len(list) == 0 {
		return 0
	}
	m := map[string]int{}
	count := 0
	for _, v := range list {
		if m[v.Name] == 0 {
			count++
		}
		m[v.Name] = 1
	}
	return count
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
