package entities

import "gorm.io/gorm"

type Criteria struct {
	gorm.Model
	Name           string
	Value          float32
	CriteriaListID uint
}

type CriteriaList struct {
	gorm.Model
	List []Criteria
}
