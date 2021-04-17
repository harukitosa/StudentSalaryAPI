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

// ToSalaryMap 各要素ごとにMapする
func (j *JobSalaryService) ToSalaryMap(list []domain.JobSalary) map[string]int {
	m := map[string]int{}

	for _, v := range list {
		s := v.Salary
		switch {
		case s < 1000:
			m["0"]++
			break
		case 1000 <= s && s < 2000:
			m["1000"]++
			break
		case 2000 <= s && s < 3000:
			m["2000"]++
			break
		case 3000 <= s && s < 4000:
			m["3000"]++
			break
		case 4000 <= s && s < 5000:
			m["4000"]++
			break
		case 5000 <= s && s < 6000:
			m["5000"]++
			break
		case 6000 <= s && s < 7000:
			m["6000"]++
			break
		case 7000 <= s && s < 8000:
			m["7000"]++
			break
		case 8000 <= s && s < 9000:
			m["8000"]++
			break
		case 9000 <= s && s < 10000:
			m["9000"]++
			break
		default:
			m["10000"]++
			break
		}
	}
	return m
}
