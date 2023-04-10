package routers

import (
	"image-server/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20
	r.POST("/image", controllers.UploadImage)

	return r
}
