package graph

import (
	"studentSalaryAPI/graph/model"
	"studentSalaryAPI/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Review   repository.ReviewRepository
	Workdata repository.WorkDataRepository
	Company  repository.CompanyRepository
	// とりあえずここに
	Blog []*model.Blog
}
