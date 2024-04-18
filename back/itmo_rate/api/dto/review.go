package dto

import (
	"itmo_rate/DB/entities"
	"itmo_rate/util"
	"time"

	"gorm.io/gorm"
)

type ReviewDTO struct {
	ID        uint          `json:"id"`
	Lecturers []string      `json:"lecturers"`
	Teachers  []string      `json:"teachers"`
	Subject   string        `json:"subject"`
	Text      string        `json:"text"`
	Author    string        `json:"author"`
	Created   time.Time     `json:"created"`
	Criteria  []CriteriaDTO `json:"criteria"`
}

func ReviewDTOFromReview(db *gorm.DB, review *entities.Review) (ret ReviewDTO, err error) {
	err = db.Preload("CriteriaList.List").First(review).Error
	if err != nil {
		return
	}

	println(len(review.CriteriaList.List))

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
	ID      uint      `json:"id"`
	Rating  []float32 `json:"rating"`
	Text    string    `json:"text"`
	Created time.Time `json:"created"`
	Author  string    `json:"author"`
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
