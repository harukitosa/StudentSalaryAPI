package domain

type Company struct {
	Name     string
	Max      int
	Min      int
	Count    int
	WorkData []WorkData
}

type CompanyRepository interface {
	SelectByTop() ([]Company, error)
	SelectByName(name string) (*Company, error)
}
