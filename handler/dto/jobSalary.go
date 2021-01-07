package dto

// JobSalaryBody is user model
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
