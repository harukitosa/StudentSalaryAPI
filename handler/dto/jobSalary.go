package dto

// JobSalaryBody is user domain
type JobSalaryBody struct {
	CreateDataJS string `json:"createDataJS"`
	Detail       string `json:"detail"`
	Experience   string `json:"experience"`
	Name         string `json:"name"`
	Salary       int    `json:"salary"`
	Term         string `json:"term"`
	Type         string `json:"type"`
	WorkDays     string `json:"work_days"`
	WorkType     string `json:"work_type"`
}

// ExportJobSalaryBody is
// Salaryはstringで保存していたのでstringで受け取る
type ExportJobSalaryBody struct {
	CreateDataJS int    `json:"create_date"`
	Detail       string `json:"detail"`
	Experience   string `json:"experience"`
	Name         string `json:"name"`
	Salary       string `json:"salary"`
	Term         string `json:"term"`
	Type         string `json:"type"`
	WorkDays     string `json:"work_days"`
	WorkType     string `json:"work_type"`
}

// JobSalaryStatistics is
type JobSalaryStatistics struct {
	Count        int `json:"count"`
	CompanyCount int `json:"company_count"`
	Avg          int `json:"avg"`
	Mid          int `json:"mid"`
}

// Map
type Map struct {
	M   map[string]int `json:"map"`
	Len int            `json:"len"`
}
