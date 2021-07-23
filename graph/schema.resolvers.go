package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"studentSalaryAPI/domain"
	"studentSalaryAPI/graph/generated"
	"studentSalaryAPI/graph/model"
)

func (r *mutationResolver) CreateWorkData(ctx context.Context, input model.NewWorkData) (*model.WorkData, error) {
	log.Printf("%+v\n", input)
	workdata := domain.WorkData{
		Name:         *input.Name,
		Salary:       input.Salary,
		CreateDataJS: *input.CreateDataJs,
		Detail:       *input.Detail,
		Experience:   *input.Experience,
		IsShow:       *input.IsShow,
		Term:         *input.Term,
		Type:         *input.Type,
		WorkDays:     *input.Workdays,
		WorkType:     *input.WorkType,
	}
	id, err := r.Resolver.Workdata.Insert(workdata)
	if err != nil {
		return nil, err
	}
	response := &model.WorkData{
		ID:           fmt.Sprint(id),
		Name:         *input.Name,
		Salary:       input.Salary,
		CreateDataJs: input.CreateDataJs,
		Detail:       input.Detail,
		Experience:   input.Experience,
		IsShow:       input.IsShow,
		Term:         input.Term,
		Type:         input.Type,
		Workdays:     input.Workdays,
		WorkType:     input.WorkType,
	}
	return response, nil
}

func (r *queryResolver) Workdata(ctx context.Context) ([]*model.WorkData, error) {
	workdata, err := r.Resolver.Workdata.SelectAll()
	if err != nil {
		return nil, err
	}
	var dto []*model.WorkData
	for _, v := range workdata {
		dto = append(dto, &model.WorkData{
			ID:           fmt.Sprint(v.ID),
			Name:         v.Name,
			Salary:       v.Salary,
			CreateDataJs: &v.CreateDataJS,
			Detail:       &v.Detail,
			Experience:   &v.Experience,
			IsShow:       &v.IsShow,
			Term:         &v.Term,
			Type:         &v.Type,
			Workdays:     &v.WorkDays,
			WorkType:     &v.WorkType,
		})
	}
	return dto, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
