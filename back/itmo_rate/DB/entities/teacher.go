package entities

import "gorm.io/gorm"

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
	TeacherID uint   `gorm:"primaryKey"`
	ReviewID  uint   `gorm:"primaryKey"`
	Role      string `gorm:"primaryKey"`
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

func (teacher *Teacher) AddReview(db *gorm.DB, review *Review, role string) error {
	return db.Create(NewTeacherReviewRel(teacher, review, role)).Error
}
