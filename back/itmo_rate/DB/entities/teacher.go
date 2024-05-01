package entities

import (
	"itmo_rate/util"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name             string
	ReviewsCount     uint
	MeanCriteriaList CriteriaList
	Reviews          []Review  `gorm:"many2many:teacher_review_rel;"`
	Subjects         []Subject `gorm:"many2many:subject_teacher_rel;"`
}

func NewTeacher(name string) Teacher {
	return Teacher{
		Name:             name,
		ReviewsCount:     0,
		MeanCriteriaList: NewDefaultCriteriaList(0, 0, 0, 0, 0).ToCriteriaList(),
	}
}

type TeacherReviewRel struct {
	gorm.Model
	TeacherID uint
	ReviewID  uint
	Role      string
}

func (TeacherReviewRel) TableName() string {
	return "teacher_review_rel"
}

func NewTeacherReviewRel(teacher *Teacher, review *Review, role string) TeacherReviewRel {
	return TeacherReviewRel{
		TeacherID: teacher.ID,
		ReviewID:  review.ID,
		Role:      role,
	}
}

func (teacher *Teacher) AddReview(db *gorm.DB, review *Review, role string) (err error) {
	err = db.Preload("MeanCriteriaList.List").First(teacher).Error
	if err != nil {
		return
	}
	util.Apply(
		teacher.MeanCriteriaList.List,
		func(criteria *Criteria) {
			for _, reviewCriteria := range review.CriteriaList.List {
				if criteria.Name == reviewCriteria.Name {
					criteria.Value = (criteria.Value*float32(teacher.ReviewsCount) + reviewCriteria.Value) / float32(teacher.ReviewsCount+1)
				}
			}
		},
	)
	teacher.ReviewsCount++
	rel := NewTeacherReviewRel(teacher, review, role)
	return db.Create(&rel).Error
}
