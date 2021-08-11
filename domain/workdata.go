package domain

import (
	"fmt"
	"sort"
)

// domain
type WorkData struct {
	ID                int
	CreateDataJS      string
	Detail            string
	Experience        string
	IsShow            bool
	Name              string
	salary            salary
	Term              string
	engineeringDomain engineeringDomain
	WorkDays          string
	contractType      contractType
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

func (w *WorkData) GetContractType() *contractType {
	return &w.contractType
}

func (w *WorkData) GetEnginneringDomain() *engineeringDomain {
	return &w.engineeringDomain
}

//  VO: 契約種別
type contractType string

const (
	INTERN      = "インターン"
	OUTSOURCING = "業務委託"
	PARTTIME    = "アルバイト"
	NONE        = "なし"
)

func isContentType(value string) bool {
	return value == INTERN || value == OUTSOURCING || value == PARTTIME || value == NONE
}

func newcontractType(value *string) (contractType, error) {
	if value == nil || *value == "" {
		return contractType(NONE), nil
	}
	s := *value
	if isContentType(s) {
		return contractType(s), nil
	}
	return contractType(NONE), fmt.Errorf("対応する契約種別がありません")
}

func (c *contractType) String() string {
	return string(*c)
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

// VO: エンジニアリング領域
type engineeringDomain string

const (
	IOS         = "IOS"
	ANDROID     = "Android"
	MOBILE      = "Mobile"
	WEBFRONTEND = "Web Frontend"
	FULLSTACK   = "FULLSTACK"
	BACKEND     = "Backend"
	AIML        = "AL/ML"
	INFRA       = "Infra"
	SRE         = "Site Reliability(SRE)"
	SECURITY    = "Security"
	DEVOPS      = "Devops"
	DATA        = "Data"
	NETWORKING  = "Networking"
	OTHERS      = "その他"
)

func isEngineeringDomain(v string) bool {
	return v == IOS || v == ANDROID || v == MOBILE || v == WEBFRONTEND || v == FULLSTACK || v == BACKEND || v == AIML || v == INFRA || v == SRE || v == SECURITY || v == DEVOPS || v == DATA || v == NETWORKING || v == OTHERS
}

func newengineeringDomain(value *string) (engineeringDomain, error) {
	if value == nil || *value == "" {
		return engineeringDomain(OTHERS), nil
	}
	if isEngineeringDomain(*value) {
		return engineeringDomain(*value), nil
	}
	return engineeringDomain(OTHERS), fmt.Errorf("対応するエンジニアリング領域がありません")
}

func (e *engineeringDomain) getValue() *engineeringDomain {
	return e
}

func (e *engineeringDomain) String() string {
	return string(*e)
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
	contractType, err := newcontractType(workType)
	if err != nil {
		return nil, err
	}
	engineeringDomain, err := newengineeringDomain(Type)
	if err != nil {
		return nil, err
	}
	w := WorkData{
		ID:                *id,
		Name:              convertNilString(name),
		salary:            *salary,
		CreateDataJS:      convertNilString(create_data_js),
		Detail:            convertNilString(detail),
		Experience:        convertNilString(experience),
		IsShow:            convertNilBoolean(isShow),
		Term:              convertNilString(term),
		engineeringDomain: engineeringDomain,
		WorkDays:          convertNilString(workDays),
		contractType:      contractType,
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
