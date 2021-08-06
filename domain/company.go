package domain

type Company struct {
	Name  string
	Max   int
	Min   int
	Count int
}

type CompanyRepository interface {
	SelectByTop() ([]Company, error)
}
