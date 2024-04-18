package entities

import "gorm.io/gorm"

type Faculty struct {
	gorm.Model
	Name     string
	Subjects []Subject `gorm:"many2many:subject_faculty_rel;"`
}

func NewFaculty(name string) Faculty {
	return Faculty{
		Name: name,
	}
}

func (faculty *Faculty) AddSubjects(db *gorm.DB, subjects []Subject) error {
	return db.Model(faculty).Association("Subjects").Append(subjects)
}
