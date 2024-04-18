package dto

import (
	"itmo_rate/DB/entities"
	"itmo_rate/util"
)

type TeacherDTO struct {
	ID        uint
	Name      string
	Criteria  []CriteriaDTO
	AvgRating float32
	Subjects  []string
}

func TeacherDTOFromTeacher(teacher *entities.Teacher) TeacherDTO {
	criteriaList := CriteriaDTOListFromCriteriaList(&teacher.MeanCriteriaList)
	return TeacherDTO{
		ID:       teacher.ID,
		Name:     teacher.Name,
		Criteria: criteriaList,
		AvgRating: util.Mean(
			util.Map(
				criteriaList,
				func(criteria CriteriaDTO) float32 {
					return criteria.Rating
				},
			),
		),
		Subjects: util.Map(
			teacher.Subjects,
			func(subject entities.Subject) string {
				return subject.Name
			},
		),
	}
}

func (teacher TeacherDTO) ToObjectPreviewDTO() ObjectPreviewDTO {
	return ObjectPreviewDTO{
		ID:    teacher.ID,
		Name:  teacher.Name,
		Score: teacher.AvgRating,
		Tags:  teacher.Subjects,
	}
}
