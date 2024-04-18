package entities

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name             string
	ReviewsCount     uint
	MeanCriteriaList CriteriaList
	Teachers         []Teacher `gorm:"many2many:subject_teacher_rel;"`
	Reviews          []Review
	Faculties        []Faculty `gorm:"many2many:subject_faculty_rel;"`
}

func NewSubject(name string) Subject {
	return Subject{
		Name:             name,
		ReviewsCount:     0,
		MeanCriteriaList: NewDefaultCriteriaList(0, 0, 0, 0, 0).ToCriteriaList(),
	}
}

func (subject *Subject) AddReviews(db *gorm.DB, reviews []Review) error {
	return db.Model(subject).Association("Reviews").Append(reviews)
}

func (subject *Subject) GetTeachersByRole(db *gorm.DB, role string) (teachers []Teacher, err error) {
	err = db.Model(subject).Association("Teachers").Find(&teachers, "role = ?", role)
	return
}

type SubjectTeacherRel struct {
	SubjectID uint   `gorm:"primaryKey"`
	TeacherID uint   `gorm:"primaryKey"`
	Role      string `gorm:"primaryKey"`
}

func (SubjectTeacherRel) TableName() string {
	return "subject_teacher_rel"
}

func NewSubjectTeacherRel(subject *Subject, teacher *Teacher, role string) SubjectTeacherRel {
	return SubjectTeacherRel{
		SubjectID: subject.ID,
		TeacherID: teacher.ID,
		Role:      role,
	}
}

func (subject *Subject) AddTeacher(db *gorm.DB, teacher *Teacher, role string) error {
	return db.Create(NewSubjectTeacherRel(subject, teacher, role)).Error
}
