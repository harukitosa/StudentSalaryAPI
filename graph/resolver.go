package graph

import (
	"studentSalaryAPI/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Review   repository.ReviewRepository
	Workdata repository.WorkDataRepository
	Company  repository.CompanyRepository
	Blog     repository.BlogRepository
}
