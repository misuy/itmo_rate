package dto

import (
	"itmo_rate/DB/entities"
	"itmo_rate/util"

	"gorm.io/gorm"
)

type SubjectDTO struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	Faculties []string      `json:"faculties"`
	Criteria  []CriteriaDTO `json:"criteria"`
	AvgRating float32       `json:"avg_rating"`
	Lecturers []string      `json:"lecturers"`
	Teachers  []string      `json:"teachers"`
}

func SubjectDTOFromSubject(db *gorm.DB, subject *entities.Subject) (ret SubjectDTO, err error) {
	err = db.Preload("MeanCriteriaList.List").Preload("Faculties").First(subject).Error
	if err != nil {
		return
	}

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
