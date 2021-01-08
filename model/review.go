package model

import "gorm.io/gorm"

// Review is user model
type Review struct {
	gorm.Model
	CompanyName  string
	Content      string
	CreateDateJS string
	Link         string
	Reasons      string
	Report       string
	Skill        string
	UserName     string
}
