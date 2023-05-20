package handler

import (
	"image-server/src/modules/image/services"
)

type HttpHandlerImpl struct {
	services.ImageService
}

func NewHttpHandler(imageService services.ImageService) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		ImageService: imageService,
	}
}
