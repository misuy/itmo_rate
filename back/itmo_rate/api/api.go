package api

import (
	"fmt"
	"itmo_rate/api/endpoints"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(db *gorm.DB, port uint) {
	router := gin.Default()

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

	router.Run(fmt.Sprintf(":%d", port))
}
