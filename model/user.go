package model

// User is user model
type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
