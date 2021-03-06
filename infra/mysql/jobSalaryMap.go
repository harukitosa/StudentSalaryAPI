package mysql

import (
	"studentSalaryAPI/domain"
	"studentSalaryAPI/repository"

	"gorm.io/gorm"
)

// JobSalaryMapRepositoryImpl is dependent sqlx and JobSalaryRepository
type JobSalaryMapRepositoryImpl struct {
	db *gorm.DB
}

// NewJobSalaryMapRepositoryImpl is
func NewJobSalaryMapRepositoryImpl(DB *gorm.DB) repository.JobSalaryMapRepository {
	return &JobSalaryMapRepositoryImpl{
		db: DB,
	}
}

// Select is
func (r *JobSalaryMapRepositoryImpl) Select() ([]domain.JobSalaryMap, error) {
	var list []domain.JobSalaryMap
	tx := r.db.Table("job_salaries").Select("name, max(salary) as max, min(salary) as min, count(*) as count").Group("name").Find(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return list, nil
}

// SelectByCount is
func (r *JobSalaryMapRepositoryImpl) SelectByCount() ([]domain.JobSalaryMap, error) {
	var list []domain.JobSalaryMap
	tx := r.db.Table("job_salaries").Select("name, max(salary) as max, min(salary) as min, count(*) as count").Group("name").Order("count desc").Limit(3).Find(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return list, nil
}
