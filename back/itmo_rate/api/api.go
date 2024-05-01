package api

import (
	"fmt"
	"itmo_rate/api/endpoints"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Access-Control-Allow-Credentials, Access-Control-Allow-Origin, Access-Control-Allow-Methods, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Run(db *gorm.DB, port uint) {
	router := gin.Default()
	router.Use(CORSMiddleware())
	api := router.Group("/api")
	api.GET("/user", endpoints.UserEndpoint(db))
	api.GET("/search", endpoints.SearchEndpoint(db))
	api.GET("/teachers", endpoints.TeachersEndpoint(db))
	api.GET("/subjects", endpoints.SubjectsEndpoint(db))
	api.GET("/subject/:id", endpoints.SubjectEndpoint(db))
	api.GET("/teacher/:id", endpoints.TeacherEndpoint(db))
	api.GET("/subject/:id/reviews", endpoints.SubjectReviewsEndpoint(db))
	api.GET("/teacher/:id/reviews", endpoints.TeacherReviewsEndpoint(db))
	api.GET("/review/:id", endpoints.ReviewEndpoint(db))

	api.POST("/review", endpoints.AddReviewEndpoint(db))

	router.Run(fmt.Sprintf(":%d", port))
}
