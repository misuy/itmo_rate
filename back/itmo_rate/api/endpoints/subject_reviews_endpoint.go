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

func SubjectReviewsEndpoint(db *gorm.DB) gin.HandlerFunc {
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

		var reviews []entities.Review
		err = db.Offset(offset).Limit(amount).Find(&reviews, "subject_id = ?", id).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

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
