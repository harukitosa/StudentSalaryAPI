package service

import "studentSalaryAPI/domain"

// JobSalaryService is
type JobSalaryService struct{}

// GetAvg is 平均値を返す
func (j *JobSalaryService) GetAvg(list []domain.JobSalary) int {
	avg := 0
	for _, v := range list {
		avg += v.Salary
	}
	return avg / len(list)
}

// GetMid 中央値を返す
func (j *JobSalaryService) GetMid(list []domain.JobSalary) int {
	size := len(list)
	if size%2 == 0 {
		return (list[size/2].Salary + list[size/2-1].Salary) / 2
	}
	return list[size/2].Salary
}

// GetCountByCompanyName ユニークな会社名がいくつあるのか調べる
func (j *JobSalaryService) GetCountByCompanyName(list []domain.JobSalary) int {
	m := map[string]int{}
	count := 0
	for _, v := range list {
		if m[v.Name] == 0 {
			count++
		}
		m[v.Name] = 1
	}
	return count
}
