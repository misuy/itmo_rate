package endpoints

import (
	"itmo_rate/DB/entities"
	"itmo_rate/api/dto"
	"itmo_rate/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TeacherReviewsEndpoint(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		offset := 0
		if offsett, ok := ctx.GetQuery("offset"); ok {
			if offsett, err := strconv.Atoi(offsett); err == nil {
				offset = offsett
			}
		}

		amount := 0
		if amountt, ok := ctx.GetQuery("amount"); ok {
			if amountt, err := strconv.Atoi(amountt); err == nil {
				amount = amountt
			}
		}

		teacher, err := entities.GetById[entities.Teacher](db, uint(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}

		err = db.Preload("Reviews", func(tx *gorm.DB) *gorm.DB {
			return tx.Offset(offset).Limit(amount)
		}).Find(&teacher).Error

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		reviews := teacher.Reviews

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"reviews": util.Map(
					reviews,
					func(review entities.Review) dto.ReviewPreviewDTO {
						reviewDTO, err := dto.ReviewDTOFromReview(db, &review)
						if err != nil {
							ctx.JSON(http.StatusInternalServerError, gin.H{})
							return dto.ReviewPreviewDTO{}
						}
						return reviewDTO.ToReviewPreviewDTO()
					},
				),
			},
		)
	}
}
