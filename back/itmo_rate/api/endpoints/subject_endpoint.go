package endpoints

import (
	"itmo_rate/DB/entities"
	"itmo_rate/api/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SubjectEndpoint(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		subject, err := entities.GetById[entities.Subject](db, uint(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}

		subjectDTO, err := dto.SubjectDTOFromSubject(db, &subject)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		println(subjectDTO.AvgRating, subjectDTO.Name, subjectDTO.Teachers, subjectDTO.Faculties)
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"subject": subjectDTO,
			},
		)
	}
}
