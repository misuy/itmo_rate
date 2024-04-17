package entities

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	ReviewsCount     uint
	CriteriaListID   uint
	MeanCriteriaList CriteriaList
	Teachers         []*Teacher `gorm:"many2many:subject_teacher_rel;"`
	Reviews          []*Review
	Faculties        []*Faculty `gorm:"many2many:subject_faculty_rel;"`
}
