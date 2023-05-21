package routers

import (
	"image-server/database"
	handler "image-server/src/handlers/http"
	"image-server/src/modules/image/repositories"
	"image-server/src/modules/image/services"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	db := database.GetDB()
	imageRepository := repositories.NewRepository(db)
	imageService := services.NewService(imageRepository)
	imageHandler := handler.NewHttpHandler(imageService)

	r.MaxMultipartMemory = 8 << 20
	r.Static("/image", "./uploads/image")
	r.POST("/imageUp", imageHandler.UploadImageHandler)
	r.DELETE("/image/:imageId", imageHandler.DeleteImageHandler)
	r.PUT("/image/:imageId", imageHandler.UpdateImageHandler)

	return r
}
