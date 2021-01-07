package model

import "gorm.io/gorm"

// User is user model
type User struct {
	gorm.Model
	Name string `db:"name" json:"name"`
}
