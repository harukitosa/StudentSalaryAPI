package mysql

import (
	"studentSalaryAPI/model"
	"studentSalaryAPI/repository"

	"gorm.io/gorm"
)

// JobSalaryRepositoryImpl is dependent sqlx and JobSalaryRepository
type JobSalaryRepositoryImpl struct {
	db *gorm.DB
}

// NewJobSalaryRepositoryImpl is
func NewJobSalaryRepositoryImpl(DB *gorm.DB) repository.JobSalaryRepository {
	return &JobSalaryRepositoryImpl{
		db: DB,
	}
}

// Insert is
func (r *JobSalaryRepositoryImpl) Insert(jobSalary model.JobSalary) (int, error) {
	result := r.db.Create(&jobSalary)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(jobSalary.ID), nil
}

// SelectByID is
func (r *JobSalaryRepositoryImpl) SelectByID(id int) (model.JobSalary, error) {
	var jobSalary model.JobSalary
	tx := r.db.First(&jobSalary, id)
	return jobSalary, tx.Error
}

// SelectAll is
func (r *JobSalaryRepositoryImpl) SelectAll() ([]model.JobSalary, error) {
	var jobSalaries []model.JobSalary
	result := r.db.Order("create_data_js desc").Find(&jobSalaries)
	return jobSalaries, result.Error
}

// SelectByName is
func (r *JobSalaryRepositoryImpl) SelectByName(name string) ([]model.JobSalary, error) {
	var jobSalaries []model.JobSalary
	result := r.db.Order("create_data_js desc").Where("name = ?", name).Find(&jobSalaries)
	return jobSalaries, result.Error
}
