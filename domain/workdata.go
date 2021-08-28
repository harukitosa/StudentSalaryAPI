package domain

import (
	"fmt"
	"sort"
)

// entity
type WorkData struct {
	id                workdataID
	createDate        createDate
	detail            workdetail
	experience        experience
	isShow            approval
	name              companyName
	salary            salary
	term              workterm
	engineeringDomain engineeringDomain
	workdays          workdays
	contractType      contractType
}

func (w *WorkData) GetID() *workdataID {
	return &w.id
}

func (w *WorkData) GetApprove() *approval {
	return &w.isShow
}

func (w *WorkData) GetCreateDate() *createDate {
	return &w.createDate
}

func (w *WorkData) GetExperience() *experience {
	return &w.experience
}

func (w *WorkData) GetTerm() *workterm {
	return &w.term
}

func (w *WorkData) GetWorkDetail() *workdetail {
	return &w.detail
}

func (w *WorkData) GetCompanyName() *companyName {
	return &w.name
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

func (w *WorkData) GetWorkDays() *workdays {
	return &w.workdays
}

// ENTITY: WorkDataID
type workdataID int

func newworkdataID(value *int) (workdataID, error) {
	if value == nil {
		return workdataID(0), fmt.Errorf("idの値がnilです")
	}
	return workdataID(*value), nil
}

func (w *workdataID) Int() int {
	return int(*w)
}

// VO: 経験年数
// サボってます
type experience string

func newexperience(value *string) experience {
	if value == nil || *value == "" {
		return experience("未記入")
	}
	str := *value
	return experience(str)
}

func (e *experience) String() string {
	return string(*e)
}

// VO: 承認済み
type approval bool

func newapproval(value *bool) approval {
	if value == nil {
		return approval(false)
	}
	return approval(*value)
}

func (a *approval) approve() {
	approve := approval(true)
	a = &approve
}

func (a *approval) Bool() bool {
	return bool(*a)
}

// VO: 勤務期間
type workterm string

const (
	oneday            = "1day"
	twodays           = "2days"
	week              = "1week"
	twoweeks          = "2weeks"
	threeweeks        = "3weeks"
	month             = "1month"
	twoorthreemonth   = "2~3month"
	halfyear          = "6months"
	year              = "1year"
	twoorthreeyear    = "2~3year"
	morethanthreeyear = "More than 3 years"
)

func isworkterm(value string) bool {
	return value == oneday || value == twodays || value == week || value == twoweeks || value == threeweeks || value == month || value == twoorthreemonth || value == halfyear || value == year || value == twoorthreeyear || value == morethanthreeyear
}

func newworkterm(value *string) (workterm, error) {
	if isworkterm(*value) {
		term := workterm(*value)
		return term, nil
	}
	str := workterm("未記入")
	return str, nil
}

func (w *workterm) String() string {
	return string(*w)
}

//  VO: 契約種別
type contractType string

const (
	intern      = "インターン"
	outsourcing = "業務委託"
	parttime    = "アルバイト"
	none        = "なし"
)

func isContentType(value string) bool {
	return value == intern || value == outsourcing || value == parttime || value == none
}

func newcontractType(value *string) (contractType, error) {
	if value == nil || *value == "" {
		return contractType(none), nil
	}
	s := *value
	if isContentType(s) {
		return contractType(s), nil
	}
	return contractType(none), fmt.Errorf("対応する契約種別がありません")
}

func (c *contractType) String() string {
	return string(*c)
}

// VO: 給料
type salary int

func newsalary(value *int) (*salary, error) {
	if value == nil {
		return nil, fmt.Errorf("給料の値がnilです")
	}
	if *value < 0 {
		return nil, fmt.Errorf("給料の値が負の数です")
	}
	if *value >= 100000 {
		return nil, fmt.Errorf("給料の値が想定より高いです")
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
	ios         = "iOS"
	android     = "Android"
	mobile      = "Mobile"
	webfrontend = "Web Frontend"
	fullstack   = "Fullstack"
	backend     = "Backend"
	aiml        = "AI/ML"
	infra       = "Infra"
	sre         = "Site Reliability(SRE)"
	security    = "Security"
	devops      = "Devops"
	data        = "Data"
	networking  = "Networking"
	others      = "その他"
	nonedomain  = "記載なし"
)

func isEngineeringDomain(v string) bool {
	return v == ios || v == android || v == mobile || v == webfrontend || v == fullstack || v == backend || v == aiml || v == infra || v == sre || v == security || v == devops || v == data || v == networking || v == others || v == nonedomain
}

func newengineeringDomain(value *string) (engineeringDomain, error) {
	if value == nil || *value == "" {
		return engineeringDomain(nonedomain), nil
	}
	if isEngineeringDomain(*value) {
		return engineeringDomain(*value), nil
	}
	return engineeringDomain(nonedomain), fmt.Errorf("対応するエンジニアリング領域がありません")
}

func (e *engineeringDomain) getValue() *engineeringDomain {
	return e
}

func (e *engineeringDomain) String() string {
	return string(*e)
}

// VO: 週出勤日数
// intの方がbetterだよね
type workdays string

const (
	one          = "1"
	two          = "2"
	three        = "3"
	four         = "4"
	five         = "5"
	noneworkdays = "記載なし"
)

func newworkdays(value *string) (workdays, error) {
	if value == nil || *value == "" {
		return workdays(noneworkdays), nil
	}
	if isWorkdays(*value) {
		return workdays(*value), nil
	}
	return workdays(noneworkdays), fmt.Errorf("対応する週出勤日数がありません")
}

func isWorkdays(v string) bool {
	return v == one || v == two || v == three || v == four || v == five || v == noneworkdays
}

func (w *workdays) getValue() *workdays {
	return w
}

func (w *workdays) String() string {
	return string(*w)
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
		if m[v.GetCompanyName().String()] == 0 {
			count++
		}
		m[v.GetCompanyName().String()] = 1
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
	workdays, err := newworkdays(workDays)
	if err != nil {
		return nil, err
	}
	companyName, err := newcompanyName(name)
	if err != nil {
		return nil, err
	}

	wd := newworkdetail(detail)
	workdataID, err := newworkdataID(id)
	if err != nil {
		return nil, err
	}

	e := newexperience(experience)
	wterm, err := newworkterm(term)
	if err != nil {
		return nil, err
	}
	date := newcreateDate(create_data_js)

	w := WorkData{
		id:                workdataID,
		name:              companyName,
		salary:            *salary,
		createDate:        date,
		detail:            wd,
		experience:        e,
		isShow:            newapproval(isShow),
		term:              wterm,
		engineeringDomain: engineeringDomain,
		workdays:          workdays,
		contractType:      contractType,
	}

	return &w, nil
}
