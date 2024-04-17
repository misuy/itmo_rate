package entities

import "gorm.io/gorm"

type Faculty struct {
	gorm.Model
	Name     string
	Subjects []*Subject `gorm:"many2many:subject_faculty_rel;"`
}
