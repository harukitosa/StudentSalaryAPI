package model

// JobSalaryMap is user model
// JobSalaryにうまくまとめられそう
type JobSalaryMap struct {
	Name  string `json:"name"`
	Max   int    `json:"max"`
	Min   int    `json:"min"`
	Count int    `json:"count"`
}
