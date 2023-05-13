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
	r.POST("/image", handler.UploadImage)
	r.POST("/imageUp", imageHandler.UploadImageVer2)
	r.DELETE("/image/:imageId", handler.DeleteImage)
	r.PATCH("/image/:imageId", handler.UpdateImage)

	return r
}
