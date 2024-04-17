package entities

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID         uint
	SubjectID      uint
	LecturerID     uint
	TeacherID      uint
	CriteriaListID CriteriaList
	CriteriaList   CriteriaList
	Text           string
	IsAnonymous    bool
}
