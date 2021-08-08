// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Company struct {
	Name     string      `json:"name"`
	Max      int         `json:"max"`
	Min      int         `json:"min"`
	Count    int         `json:"count"`
	Workdata []*WorkData `json:"workdata"`
	Review   []*Review   `json:"review"`
}

type NewReview struct {
	CompanyName  string  `json:"company_name"`
	Detail       *string `json:"detail"`
	Content      string  `json:"content"`
	CreateDataJs *string `json:"create_data_js"`
	Link         *string `json:"link"`
	Reasons      *string `json:"reasons"`
	Report       *string `json:"report"`
	Skill        *string `json:"skill"`
	UserName     *string `json:"user_name"`
}

type NewWorkData struct {
	CreateDataJs *string `json:"create_data_js"`
	Detail       *string `json:"detail"`
	Experience   *string `json:"experience"`
	IsShow       *bool   `json:"is_show"`
	Name         *string `json:"name"`
	Salary       int     `json:"salary"`
	Term         *string `json:"term"`
	Type         *string `json:"type"`
	Workdays     *string `json:"workdays"`
	WorkType     *string `json:"workType"`
}

type Review struct {
	ID           string  `json:"id"`
	CompanyName  *string `json:"company_name"`
	Detail       *string `json:"detail"`
	Content      *string `json:"content"`
	CreateDataJs *string `json:"create_data_js"`
	Link         *string `json:"link"`
	Reasons      *string `json:"reasons"`
	Report       *string `json:"report"`
	Skill        *string `json:"skill"`
	UserName     *string `json:"user_name"`
}

type WorkData struct {
	ID           string  `json:"id"`
	CreateDataJs *string `json:"create_data_js"`
	Detail       *string `json:"detail"`
	Experience   *string `json:"experience"`
	IsShow       *bool   `json:"is_show"`
	Name         string  `json:"name"`
	Salary       int     `json:"salary"`
	Term         *string `json:"term"`
	Type         *string `json:"type"`
	Workdays     *string `json:"workdays"`
	WorkType     *string `json:"workType"`
}

type WorkDataInfo struct {
	CompanyCount int         `json:"company_count"`
	Avarage      int         `json:"avarage"`
	Count        int         `json:"count"`
	Mid          int         `json:"mid"`
	Workdata     []*WorkData `json:"workdata"`
}
