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

func TeachersEndpoint(db *gorm.DB) gin.HandlerFunc {
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

		var teachers []entities.Teacher

		err := db.Offset(offset).Limit(amount).Find(&teachers).Error

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"teachers": util.Map(
					teachers,
					func(teacher entities.Teacher) dto.ObjectPreviewDTO {
						ret, err := dto.TeacherDTOFromTeacher(db, &teacher)
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
