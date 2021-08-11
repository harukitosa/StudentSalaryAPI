package domain

import (
	"fmt"
	"sort"
)

// domain
type WorkData struct {
	ID           int
	CreateDataJS string
	Detail       string
	Experience   string
	IsShow       bool
	Name         string
	salary       salary
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

func (w *WorkData) GetSalary() *salary {
	return &w.salary
}

// VO: 給料
type salary int

func newsalary(value *int) (*salary, error) {
	if value == nil {
		return nil, fmt.Errorf("給料の値がnullです")
	}
	if *value < 0 {
		return nil, fmt.Errorf("給料の値が負の数です")
	}
	n := salary(*value)
	return &n, nil
}

func (s *salary) getValue() *salary {
	return s
}

func (s *salary) isBiggerThan(c *salary) bool {
	return s.Int() > c.Int()
}

func (s *salary) Int() int {
	return int(*s)
}

type WorkDataService struct{}

func (s *WorkDataService) GetSalaryAvg(list []WorkData) int {
	if len(list) == 0 {
		return 0
	}
	ans := 0
	for _, v := range list {
		ans += v.salary.Int()
	}
	return ans / len(list)
}

func (s *WorkDataService) GetSalaryMid(list []WorkData) int {
	if len(list) == 0 {
		return 0
	}
	// sort.Slice(list, func(i, j int) bool { return list[i].salary < list[j].salary })
	sort.Slice(list, func(i, j int) bool {
		return list[i].salary.isBiggerThan(&list[j].salary)
	})
	i := len(list) / 2
	if len(list)%2 == 0 {
		return (list[i].salary.Int() + list[i-1].salary.Int()) / 2
	}
	return list[i].salary.getValue().Int()
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

// factory - 1
func NewWorkData(
	id *int,
	create_data_js *string,
	detail *string,
	experience *string,
	isShow *bool,
	name *string,
	s *int,
	term *string,
	Type *string,
	workDays *string,
	workType *string) (*WorkData, error) {
	salary, err := newsalary(s)
	if err != nil {
		return nil, err
	}
	w := WorkData{
		ID:           *id,
		Name:         convertNilString(name),
		salary:       *salary,
		CreateDataJS: convertNilString(create_data_js),
		Detail:       convertNilString(detail),
		Experience:   convertNilString(experience),
		IsShow:       convertNilBoolean(isShow),
		Term:         convertNilString(term),
		Type:         convertNilString(Type),
		WorkDays:     convertNilString(workDays),
		WorkType:     convertNilString(workType),
	}
	return &w, nil
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
