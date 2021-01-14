package test

import (
	"studentSalaryAPI/domain"
	"studentSalaryAPI/repository"
)

// JobSalaryRepositoryTestImpl is dependent sqlx and JobSalaryRepository
type JobSalaryRepositoryTestImpl struct{}

// NewJobSalaryRepositoryTestImpl is
func NewJobSalaryRepositoryTestImpl() repository.JobSalaryRepository {
	return &JobSalaryRepositoryTestImpl{}
}

// Insert is
func (r *JobSalaryRepositoryTestImpl) Insert(jobSalary domain.JobSalary) (int, error) {
	return int(1), nil
}

// SelectByID is
func (r *JobSalaryRepositoryTestImpl) SelectByID(id int) (domain.JobSalary, error) {
	data := domain.JobSalary{
		Name:         "なまえ",
		CreateDataJS: "000000000",
		Detail:       "詳細",
		Experience:   "経験",
		IsShow:       true,
		Salary:       2000,
		Term:         "2週間",
		Type:         "タイプ",
		WorkDays:     "インターン",
		WorkType:     "IOS",
	}
	data.ID = uint(id)
	return data, nil
}

// SelectAll is
func (r *JobSalaryRepositoryTestImpl) SelectAll() ([]domain.JobSalary, error) {
	var jobSalaries []domain.JobSalary
	for i := 0; i < 3; i++ {
		jobSalaries = append(jobSalaries, domain.JobSalary{
			Name:         "なまえ",
			CreateDataJS: "000000000",
			Detail:       "詳細",
			Experience:   "経験",
			IsShow:       true,
			Salary:       2000,
			Term:         "2週間",
			Type:         "タイプ",
			WorkDays:     "インターン",
			WorkType:     "IOS",
		})
	}
	return jobSalaries, nil
}

// SelectByName is
func (r *JobSalaryRepositoryTestImpl) SelectByName(name string) ([]domain.JobSalary, error) {
	var jobSalaries []domain.JobSalary
	for i := 0; i < 3; i++ {
		jobSalaries = append(jobSalaries, domain.JobSalary{
			Name:         "なまえ",
			CreateDataJS: "000000000",
			Detail:       "詳細",
			Experience:   "経験",
			IsShow:       true,
			Salary:       2000,
			Term:         "2週間",
			Type:         "タイプ",
			WorkDays:     "インターン",
			WorkType:     "IOS",
		})
	}
	return jobSalaries, nil
}
