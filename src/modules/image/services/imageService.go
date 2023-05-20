package services

import (
	"image-server/src/modules/image/dto"
	"image-server/src/modules/image/entities"
	"image-server/src/modules/image/repositories"
)

type service struct {
	repository repositories.RepositoryImageCommand
}

func NewService(repository repositories.RepositoryImageCommand) *service {
	service := &service{repository}

	return service
}

type ImageService interface {
	InsertImage(data dto.ImageRequestBody) (entities.Image, error)
	UpdateImage(data dto.ImageRequestBody) (entities.Image, error)
	ImageByID(ID int) (entities.Image, error)
}

func (s service) ImageByID(ID int) (entities.Image, error) {
	image, err := s.repository.ImageByID(ID)

	return image, err
}

func (s service) UpdateImage(data dto.ImageRequestBody) (entities.Image, error) {
	image := entities.Image{
		Name: data.Name,
	}

	newImage, err := s.repository.UpdateImage(image)

	return newImage, err
}

func (s service) InsertImage(data dto.ImageRequestBody) (entities.Image, error) {
	image := entities.Image{
		Name:     data.Name,
		ImageUrl: data.ImageUrl,
		UserID:   data.UserID,
	}

	newImage, err := s.repository.InsertImage(image)

	return newImage, err
}
