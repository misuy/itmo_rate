package dto

import (
	"itmo_rate/DB/entities"
	"itmo_rate/util"

	"gorm.io/gorm"
)

type TeacherDTO struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	Criteria  []CriteriaDTO `json:"criteria"`
	AvgRating float32       `json:"avg_rating"`
	Subjects  []string      `json:"subjects"`
}

func TeacherDTOFromTeacher(db *gorm.DB, teacher *entities.Teacher) (ret TeacherDTO, err error) {
	err = db.Preload("MeanCriteriaList.List").Preload("Subjects").First(teacher).Error
	if err != nil {
		return
	}
	println(len(teacher.MeanCriteriaList.List))
	criteriaList := CriteriaDTOListFromCriteriaList(&teacher.MeanCriteriaList)
	ret = TeacherDTO{
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
	return
}

func (teacher TeacherDTO) ToObjectPreviewDTO() ObjectPreviewDTO {
	return ObjectPreviewDTO{
		ID:    teacher.ID,
		Name:  teacher.Name,
		Score: teacher.AvgRating,
		Tags:  teacher.Subjects,
	}
}
