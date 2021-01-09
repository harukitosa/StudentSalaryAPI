package domain

import "gorm.io/gorm"

// JobSalary is user domain
type JobSalary struct {
	gorm.Model
	CreateDataJS string `json:"create_data_js"`
	Detail       string `json:"detail"`
	Experience   string `json:"experience"`
	IsShow       bool   `json:"is_show"`
	Name         string `json:"name"`
	Salary       int    `json:"salary"`
	Term         string `json:"term"`
	Type         string `json:"type"`
	WorkDays     string `json:"work_days"`
	WorkType     string `json:"work_type"`
}
