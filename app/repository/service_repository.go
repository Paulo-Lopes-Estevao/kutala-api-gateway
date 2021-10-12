package repository

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/jinzhu/gorm"
)

type ServiceRepositoryInterface interface {
	InsertService(service *entities.Service) (*entities.Service, error)
	FindUsernameService(username string, service *entities.Service) (*entities.Service, error)
	FindIdService(id string, service *entities.Service) (*entities.Service, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepositoryInterface {
	return &serviceRepository{db}
}

func (repository *serviceRepository) InsertService(service *entities.Service) (*entities.Service, error) {
	err := repository.db.Table("service").Create(service).Error

	if err != nil {
		return nil, err
	}

	return service, nil
}

func (repository *serviceRepository) FindUsernameService(username string, service *entities.Service) (*entities.Service, error) {
	err := repository.db.Table("service").Find(service, "username = ?", username).Error

	if err != nil {
		return nil, err
	}

	return service, nil

}

func (repository *serviceRepository) FindIdService(id string, service *entities.Service) (*entities.Service, error) {
	err := repository.db.Table("service").Find(service, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return service, nil

}
