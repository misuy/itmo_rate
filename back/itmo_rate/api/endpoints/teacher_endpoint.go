package endpoints

import (
	"itmo_rate/DB/entities"
	"itmo_rate/api/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TeacherEndpoint(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		teacher, err := entities.GetById[entities.Teacher](db, uint(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}

		teacherDTO, err := dto.TeacherDTOFromTeacher(db, &teacher)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"teacher": teacherDTO,
			},
		)
	}
}
