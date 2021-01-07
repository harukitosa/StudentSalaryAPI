package model

import "gorm.io/gorm"

// JobSalary is user model
type JobSalary struct {
	gorm.Model
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
