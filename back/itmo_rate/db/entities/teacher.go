package entities

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name             string
	ReviewsCount     uint
	CriteriaListID   uint
	MeanCriteriaList CriteriaList
	Subjects         []*Subject `gorm:"many2many:subject_teacher_rel;"`
	Reviews          []*Review
}
