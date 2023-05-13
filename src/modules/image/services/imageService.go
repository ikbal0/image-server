package services

import (
	"image-server/src/modules/image/dto"
	"image-server/src/modules/image/entities"
	"image-server/src/modules/image/repositories"
)

type ImageService interface {
	InsertImage(data dto.ImageRequestBody) (entities.Image, error)
}

type service struct {
	repository repositories.RepositoryImageCommand
}

func NewService(repository repositories.RepositoryImageCommand) *service {
	service := &service{repository}

	return service
}

func (s service) InsertImage(data dto.ImageRequestBody) (entities.Image, error) {
	image := entities.Image{
		Name:     data.Name,
		ImageUrl: data.ImageUrl,
		UserID:   data.UserID,
	}

	newImage, err := s.repository.InsertImage(image)

	if err != nil {
		return newImage, err
	}

	return newImage, nil
}
