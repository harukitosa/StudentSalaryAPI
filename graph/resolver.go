package graph

import "studentSalaryAPI/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Review   domain.ReviewRepository
	Workdata domain.WorkDataRepository
}
