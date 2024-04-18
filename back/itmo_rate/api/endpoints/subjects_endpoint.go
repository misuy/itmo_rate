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

func SubjectsEndpoint(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		var subjects []entities.Subject

		err := db.Offset(offset).Limit(amount).Find(&subjects).Error

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"subjects": util.Map(
					subjects,
					func(subject entities.Subject) dto.ObjectPreviewDTO {
						ret, err := dto.SubjectDTOFromSubject(db, &subject)
						if err != nil {
							ctx.JSON(http.StatusInternalServerError, gin.H{})
							return dto.ObjectPreviewDTO{}
						}
						return ret.ToObjectPreviewDTO()
					},
				),
			},
		)
	}
}
