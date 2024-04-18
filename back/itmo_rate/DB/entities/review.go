package entities

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID       *uint
	SubjectID    *uint
	Teachers     []Teacher `gorm:"many2many:teacher_review_rel;"`
	CriteriaList CriteriaList
	Text         string
	IsAnonymous  bool
}

func NewReview(list CriteriaList, text string, isAnonymous bool) Review {
	return Review{
		CriteriaList: list,
		Text:         text,
		IsAnonymous:  isAnonymous,
	}
}

func (review *Review) GetTeachersByRole(db *gorm.DB, role string) (teachers []Teacher, err error) {
	err = db.Model(review).Association("Teachers").Find(&teachers, "role = ?", role)
	return
}
