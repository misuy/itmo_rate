package dto

import (
	"itmo_rate/DB/entities"
	"itmo_rate/util"

	"gorm.io/gorm"
)

type SubjectDTO struct {
	ID        uint
	Name      string
	Faculties []string
	Criteria  []CriteriaDTO
	AvgRating float32
	Lecturers []string
	Teachers  []string
}

func SubjectDTOFromSubject(db *gorm.DB, subject *entities.Subject) (ret SubjectDTO, err error) {
	criteriaList := CriteriaDTOListFromCriteriaList(&subject.MeanCriteriaList)
	teachers, err := subject.GetTeachersByRole(db, "teacher")
	if err != nil {
		return
	}
	lecturers, err := subject.GetTeachersByRole(db, "lecturer")
	if err != nil {
		return
	}
	ret = SubjectDTO{
		ID:   subject.ID,
		Name: subject.Name,
		Faculties: util.Map(
			subject.Faculties,
			func(faculty entities.Faculty) string {
				return faculty.Name
			},
		),
		Criteria: criteriaList,
		AvgRating: util.Mean(
			util.Map(
				criteriaList,
				func(criteria CriteriaDTO) float32 {
					return criteria.Rating
				},
			),
		),
		Lecturers: util.Map(
			lecturers,
			func(lecturer entities.Teacher) string {
				return lecturer.Name
			},
		),
		Teachers: util.Map(
			teachers,
			func(teacher entities.Teacher) string {
				return teacher.Name
			},
		),
	}
	return
}

func (subject SubjectDTO) ToObjectPreviewDTO() ObjectPreviewDTO {
	return ObjectPreviewDTO{
		ID:    subject.ID,
		Name:  subject.Name,
		Tags:  subject.Faculties,
		Score: subject.AvgRating,
	}
}
