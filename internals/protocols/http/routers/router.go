package routers

import (
	handler "image-server/src/handlers/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20
	r.Static("/image", "./image")
	r.POST("/image", handler.UploadImage)
	r.DELETE("/image/:imageId", handler.DeleteImage)
	r.PATCH("/image/:imageId", handler.UpdateImage)

	return r
}
