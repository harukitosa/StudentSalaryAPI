package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"studentSalaryAPI/domain"
	"studentSalaryAPI/graph/generated"
	"studentSalaryAPI/graph/model"
)

func (r *mutationResolver) CreateWorkData(ctx context.Context, input model.NewWorkData) (*model.WorkData, error) {
	workdata := domain.NewWorkData(
		input.CreateDataJs,
		input.Detail,
		input.Experience,
		input.IsShow,
		input.Name,
		&input.Salary,
		input.Term,
		input.Type,
		input.Workdays,
		input.WorkType)
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

func (r *mutationResolver) CreateReview(ctx context.Context, input model.NewReview) (*model.Review, error) {
	review := domain.NewReview(
		&input.CompanyName,
		input.Detail,
		&input.Content,
		input.CreateDataJs,
		input.Link,
		input.Reasons,
		input.Report,
		input.Skill,
		input.UserName,
	)
	id, err := r.Resolver.Review.Insert(review)
	if err != nil {
		return nil, err
	}
	response := model.Review{
		ID:           fmt.Sprint(id),
		CompanyName:  &review.CompanyName,
		Detail:       &review.Detail,
		Content:      &review.Content,
		CreateDataJs: &review.CreateDateJS,
		Link:         &review.Link,
		Reasons:      &review.Reasons,
		Report:       &review.Report,
		Skill:        &review.Skill,
		UserName:     &review.UserName,
	}
	return &response, nil
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

func (r *queryResolver) Review(ctx context.Context) ([]*model.Review, error) {
	reviews, err := r.Resolver.Review.SelectAll()
	if err != nil {
		return nil, err
	}
	var dto []*model.Review
	for _, v := range reviews {
		dto = append(dto, &model.Review{
			ID:           fmt.Sprint(v.ID),
			CompanyName:  &v.CompanyName,
			Detail:       &v.Detail,
			Content:      &v.Content,
			CreateDataJs: &v.CreateDateJS,
			Link:         &v.Link,
			Reasons:      &v.Reasons,
			Report:       &v.Report,
			Skill:        &v.Skill,
			UserName:     &v.UserName,
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