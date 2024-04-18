package endpoints

import (
	"fmt"
	"itmo_rate/DB/entities"
	"itmo_rate/api/dto"
	"itmo_rate/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const MAX_RESULTS = 20

func SearchEndpoint(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error

		searchType, ok := ctx.GetQuery("type")
		if !ok {
			searchType = "both"
		}

		text, ok := ctx.GetQuery("text")
		if !ok {
			text = ""
		}

		var subjects []entities.Subject
		var teachers []entities.Teacher

		switch searchType {
		case "teachers":
			err = db.Limit(MAX_RESULTS).Find(&teachers, "name like ?", fmt.Sprintf("%%%s%%", text)).Error
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		case "subjects":
			err = db.Limit(MAX_RESULTS).Find(&subjects, "name like ?", fmt.Sprintf("%%%s%%", text)).Error
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		case "both":
			err = db.Limit(MAX_RESULTS).Find(&teachers, "name like ?", fmt.Sprintf("%%%s%%", text)).Error
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
			err = db.Limit(MAX_RESULTS).Find(&subjects, "name like ?", fmt.Sprintf("%%%s%%", text)).Error
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{})
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
