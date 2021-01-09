package domain

import "gorm.io/gorm"

// User is user domain
type User struct {
	gorm.Model
	Name string `db:"name" json:"name"`
}
