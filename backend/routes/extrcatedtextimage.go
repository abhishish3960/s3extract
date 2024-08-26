// routes/routes.go
package routes

import (
	"backend/controllers"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	s3Service := services.NewS3Service("extractedtextimage")
	contentController := controllers.NewContentController(s3Service)

	router.GET("/api/contents", contentController.GetContents)
}
