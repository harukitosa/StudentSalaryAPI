package handler

import (
	"net/http"
	"studentSalaryAPI/domain"

	"github.com/labstack/echo/v4"
)

type workdataHandler struct {
	repository domain.WorkDataRepository
}

func NewWorkDataHandler(repository domain.WorkDataRepository) *workdataHandler {
	return &workdataHandler{repository: repository}
}

func (r *workdataHandler) GetReview(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return r.GetAllReview(c)
	}
	return r.GetReviewByName(c)
}

func (r *workdataHandler) GetAllReview(c echo.Context) error {
	data, err := r.repository.SelectAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	body := make([]workDataBody, len(data))
	for i, v := range data {
		body[i] = createWorkDataBody(v)
	}
	return c.JSON(http.StatusOK, body)
}

func (r *workdataHandler) GetReviewByName(c echo.Context) error {
	name := c.QueryParam("name")
	data, err := r.repository.SelectByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	body := make([]workDataBody, len(data))
	for i, v := range data {
		body[i] = createWorkDataBody(v)
	}
	return c.JSON(http.StatusOK, body)
}

func (r *workdataHandler) CreateWorkData(c echo.Context) error {
	workdata := new(workDataBody)
	err := c.Bind(workdata)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	id, err := r.repository.Insert(domain.WorkData{
		CreateDataJS: workdata.CreateDataJS,
		Detail:       workdata.Detail,
		Experience:   workdata.Experience,
		IsShow:       workdata.IsShow,
		Name:         workdata.Name,
		Salary:       workdata.Salary,
		Term:         workdata.Term,
		Type:         workdata.Type,
		WorkDays:     workdata.WorkDays,
		WorkType:     workdata.WorkType,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, id)
}
func (r *workdataHandler) GetAggregateWorkData(c echo.Context) error {
	data, err := r.repository.SelectAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, aggregateWorkData{
		Avg:          getAvg(data),
		Mid:          getMid(data),
		CompanyCount: getCountByCompanyName(data),
		Count:        len(data),
	})
}

func getAvg(list []domain.WorkData) int {
	avg := 0
	for _, v := range list {
		avg += v.Salary
	}
	return avg / len(list)
}

func getMid(list []domain.WorkData) int {
	size := len(list)
	if size%2 == 0 {
		return (list[size/2].Salary + list[size/2-1].Salary) / 2
	}
	return list[size/2].Salary
}

func getCountByCompanyName(list []domain.WorkData) int {
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

func toSalaryMap(list []domain.WorkData) map[string]int {
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

type workDataBody struct {
	ID           int    `json:"id"`
	CreateDataJS string `json:"create_data_js"`
	Detail       string `json:"detail"`
	Experience   string `json:"experience"`
	IsShow       bool   `json:"is_show"`
	Name         string `json:"name"`
	Salary       int    `json:"salary"`
	Term         string `json:"term"`
	Type         string `json:"type"`
	WorkDays     string `json:"work_days"`
	WorkType     string `json:"work_type"`
}

func createWorkDataBody(workdata domain.WorkData) workDataBody {
	return workDataBody{
		ID:           workdata.ID,
		CreateDataJS: workdata.CreateDataJS,
		Detail:       workdata.Detail,
		Experience:   workdata.Experience,
		IsShow:       workdata.IsShow,
		Name:         workdata.Name,
		Salary:       workdata.Salary,
		Term:         workdata.Term,
		Type:         workdata.Type,
		WorkDays:     workdata.WorkDays,
		WorkType:     workdata.WorkType,
	}
}

type aggregateWorkData struct {
	Count        int `json:"count"`
	CompanyCount int `json:"company_count"`
	Avg          int `json:"avg"`
	Mid          int `json:"mid"`
}
