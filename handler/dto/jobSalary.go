package dto

// JobSalaryBody is user model
// Salaryはstringで保存していたので一時てきにstringで受け取る
type JobSalaryBody struct {
	CreateDataJS string `json:"createDataJS"`
	Detail       string `json:"detail"`
	Experience   string `json:"experience"`
	Name         string `json:"name"`
	Salary       int    `json:"salary"`
	Term         string `json:"term"`
	Type         string `json:"type"`
	WorkDays     string `json:"workDays"`
	WorkType     string `json:"workType"`
}

// ExportJobSalaryBody is
type ExportJobSalaryBody struct {
	CreateDataJS string `json:"createDataJS"`
	Detail       string `json:"detail"`
	Experience   string `json:"experience"`
	Name         string `json:"name"`
	Salary       string `json:"salary"`
	Term         string `json:"term"`
	Type         string `json:"type"`
	WorkDays     string `json:"workDays"`
	WorkType     string `json:"workType"`
}
