package application

import (
	"studentSalaryAPI/domain"
	"studentSalaryAPI/infra/test"
	"testing"
)

func TestInsert(t *testing.T) {
	jobSalaryRepository := test.NewJobSalaryRepositoryTestImpl()
	jobSalaryApplication := NewJobSalaryApplication(jobSalaryRepository)
	patterns := []struct {
		a        domain.JobSalary
		expected int
	}{
		{domain.JobSalary{
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
		}, 1},
	}

	for idx, pattern := range patterns {
		actual, _ := jobSalaryApplication.Insert(pattern.a)
		if pattern.expected != actual {
			t.Errorf("pattern %d: want %d, actual %d", idx, pattern.expected, actual)
		}
	}
}

func TestGetAll(t *testing.T) {
	jobSalaryRepository := test.NewJobSalaryRepositoryTestImpl()
	jobSalaryApplication := NewJobSalaryApplication(jobSalaryRepository)

	// 3個の配列が帰ってくるかどうかを確かめる
	patterns := []struct {
		expected int
	}{
		{3},
	}

	for idx, pattern := range patterns {
		actual, _ := jobSalaryApplication.GetAll()
		if pattern.expected != len(actual) {
			t.Errorf("pattern %d: want %d, actual %d", idx, pattern.expected, len(actual))
		}
	}
}
