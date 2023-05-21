package services

import "image-server/src/modules/image/repositories"

type service struct {
	repository repositories.RepositoryImageCommand
}

func NewService(repository repositories.RepositoryImageCommand) *service {
	service := &service{repository}

	return service
}
