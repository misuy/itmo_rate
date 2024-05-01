package endpoints

import (
	"itmo_rate/api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddReviewEndpoint(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newReview dto.NewReviewDTO
		if err := ctx.ShouldBindJSON(&newReview); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		if err := newReview.AddToDB(db); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}
