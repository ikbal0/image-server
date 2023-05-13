package handler

import (
	"image-server/src/modules/image/services"

	"github.com/gin-gonic/gin"
)

type HttpHandlerImpl struct {
	services.ImageService
}

func NewHttpHandler(imageService services.ImageService) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		ImageService: imageService,
	}
}

func (h *HttpHandlerImpl) Router() {
	r := gin.Default()

	r.POST("/imageUp", h.UploadImageVer2)
}
