package endpoints

import (
	"itmo_rate/DB/entities"
	"itmo_rate/api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const USER_ID = 1

func UserEndpoint(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := entities.GetById[entities.User](db, USER_ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}
		ctx.JSON(http.StatusOK, dto.UserDTOFromUser(&user))
	}
}
