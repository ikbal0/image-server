package services

import (
	"image-server/src/modules/image/entities"
	"image-server/src/modules/image/repositories"

	"github.com/gin-gonic/gin"
)

type ImageService interface {
	InsertImage(ctx *gin.Context) (entities.Image, error)
}

type service struct {
	repository repositories.RepositoryImageCommand
}

func NewService(repository repositories.RepositoryImageCommand) *service {
	service := &service{repository}

	return service
}

func (s service) InsertImage(ctx *gin.Context) (entities.Image, error) {
	newImage, err := s.repository.InsertImage(ctx)

	return newImage, err
}
