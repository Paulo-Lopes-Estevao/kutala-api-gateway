package repository

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/jinzhu/gorm"
)

type ServiceRepositoryInterface interface {
	InsertService(service *entities.Service) (*entities.Service, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepositoryInterface {
	return &serviceRepository{db}
}

func (repository serviceRepository) InsertService(service *entities.Service) (*entities.Service, error) {
	err := repository.db.Table("service").Create(service).Error

	if err != nil {
		return nil, err
	}

	return service, nil
}
