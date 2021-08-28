package domain

import (
	"fmt"
)

type Company struct {
	Name     string
	Max      int
	Min      int
	Count    int
	WorkData []WorkData
}

// VO: 企業名
type companyName string

func newcompanyName(value *string) (companyName, error) {
	if value == nil || *value == "" {
		return companyName(""), fmt.Errorf("企業名が空です")
	}
	s := *value
	return companyName(s), nil
}

func (c companyName) String() string {
	return string(c)
}
