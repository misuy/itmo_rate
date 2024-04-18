package dto

import (
	"itmo_rate/DB/entities"
	"itmo_rate/util"
	"time"

	"gorm.io/gorm"
)

type ReviewDTO struct {
	ID        uint
	Lecturers []string
	Teachers  []string
	Subject   string
	Text      string
	Author    string
	Created   time.Time
	Criteria  []CriteriaDTO
}

func ReviewDTOFromReview(db *gorm.DB, review *entities.Review) (ret ReviewDTO, err error) {
	lecturers, err := review.GetTeachersByRole(db, "lecturer")
	if err != nil {
		return
	}
	teachers, err := review.GetTeachersByRole(db, "teacher")
	if err != nil {
		return
	}
	subject, err := entities.GetById[entities.Subject](db, *review.SubjectID)
	if err != nil {
		return
	}
	author, err := entities.GetById[entities.User](db, *review.UserID)
	if err != nil {
		return
	}

	ret = ReviewDTO{
		ID: review.ID,
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
		Subject:  subject.Name,
		Text:     review.Text,
		Author:   author.Name,
		Created:  review.CreatedAt,
		Criteria: CriteriaDTOListFromCriteriaList(&review.CriteriaList),
	}
	return
}

type ReviewPreviewDTO struct {
	ID      uint
	Rating  []float32
	Text    string
	Created time.Time
	Author  string
}

func (review ReviewDTO) ToReviewPreviewDTO() ReviewPreviewDTO {
	return ReviewPreviewDTO{
		ID: review.ID,
		Rating: util.Map(
			review.Criteria,
			func(criteria CriteriaDTO) float32 {
				return criteria.Rating
			},
		),
		Text:    review.Text,
		Created: review.Created,
		Author:  review.Author,
	}
}
