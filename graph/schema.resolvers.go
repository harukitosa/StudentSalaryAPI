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
	workdata, err := domain.NewWorkData(
		nil,
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
	if err != nil {
		return nil, err
	}
	id, err := r.Resolver.Workdata.Insert(*workdata)
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

func convertGraphqlModel(w domain.WorkData) *model.WorkData {
	v := w
	contractType := v.GetContractType().String()
	engineringDomain := v.GetEnginneringDomain().String()
	workdays := v.GetWorkDays().String()
	workdetail := v.GetWorkDetail().String()
	createdate := v.GetCreateDate().String()
	experice := v.GetExperience().String()
	isShow := v.GetApprove().Bool()
	term := v.GetTerm().String()
	return &model.WorkData{
		ID:           fmt.Sprint(v.GetID().Int()),
		Name:         v.GetCompanyName().String(),
		Salary:       v.GetSalary().Int(),
		CreateDataJs: &createdate,
		Detail:       &workdetail,
		Experience:   &experice,
		IsShow:       &isShow,
		Term:         &term,
		Type:         &engineringDomain,
		Workdays:     &workdays,
		WorkType:     &contractType,
	}
}

func (r *queryResolver) Workdatainfo(ctx context.Context) (*model.WorkDataInfo, error) {
	workdata, err := r.Resolver.Workdata.SelectAll()
	if err != nil {
		return nil, err
	}
	var dto []*model.WorkData
	for _, w := range workdata {
		dto = append(dto, convertGraphqlModel(w))
	}

	workdataService := domain.WorkDataService{}
	return &model.WorkDataInfo{
		Avarage:      workdataService.GetSalaryAvg(workdata),
		Mid:          workdataService.GetSalaryMid(workdata),
		Count:        len(workdata),
		CompanyCount: workdataService.GetCountByCompanyName(workdata),
		Workdata:     dto,
	}, nil
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

func (r *queryResolver) Newreview(ctx context.Context) ([]*model.Review, error) {
	reviews, err := r.Resolver.Review.GetNewReview()
	if err != nil {
		return nil, err
	}
	var dto []*model.Review
	tmp := append([]domain.Review{}, reviews...)
	for _, s := range tmp {
		var v domain.Review
		v = s
		r := model.Review{
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
		}
		dto = append(dto, &r)
	}
	return dto, nil
}

func (r *queryResolver) Topcompany(ctx context.Context) ([]*model.Company, error) {
	company, err := r.Resolver.Company.SelectByTop()
	if err != nil {
		return nil, err
	}
	var dto []*model.Company
	for _, v := range company {
		dto = append(dto, &model.Company{
			Name:  v.Name,
			Count: v.Count,
			Max:   v.Max,
			Min:   v.Min,
		})
	}
	return dto, nil
}

func (r *queryResolver) Company(ctx context.Context, name *string) ([]*model.Company, error) {
	if name == nil {
		company, err := r.Resolver.Company.Select()
		if err != nil {
			return nil, err
		}

		workdata, err := r.Resolver.Workdata.SelectAll()
		var workdatalist []*model.WorkData
		for _, w := range workdata {
			workdatalist = append(workdatalist, convertGraphqlModel(w))
		}

		review, err := r.Resolver.Review.SelectAll()
		if err != nil {
			return nil, err
		}
		var reviews []*model.Review
		for _, re := range review {
			v := re
			reviews = append(reviews, &model.Review{
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
		var list []*model.Company
		for _, v := range company {
			var workdata []*model.WorkData
			var reviews []*model.Review
			for _, w := range workdatalist {
				if v.Name == w.Name {
					workdata = append(workdata, w)
				}
			}
			for _, r := range reviews {
				if v.Name == *r.CompanyName {
					reviews = append(reviews, r)
				}
			}
			list = append(list, &model.Company{
				Name:     v.Name,
				Count:    v.Count,
				Max:      v.Max,
				Min:      v.Min,
				Workdata: workdata,
				Review:   reviews,
			})
		}
		return list, nil
	} else {
		company, err := r.Resolver.Company.SelectByName(*name)
		if err != nil {
			return nil, err
		}
		workdata, err := r.Resolver.Workdata.SelectByName(*name)
		var dto []*model.WorkData
		for _, w := range workdata {
			dto = append(dto, convertGraphqlModel(w))
		}

		review, err := r.Resolver.Review.SelectByName(*name)
		if err != nil {
			return nil, err
		}
		var reviews []*model.Review
		for _, re := range review {
			v := re
			reviews = append(reviews, &model.Review{
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
		var list []*model.Company
		list = append(list, &model.Company{Max: company.Max, Min: company.Min, Count: company.Count, Name: company.Name, Workdata: dto, Review: reviews})
		return list, nil
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
