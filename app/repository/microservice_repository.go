package repository

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/jinzhu/gorm"
)

type MicroserviceRepositoryInterface interface {
	InsertMicroservice(microservice *entities.Microservice) (*entities.Microservice, error)
	//FindMicroservice(name string, username string, password string) (bool, error)
}

type microserviceRepository struct {
	db *gorm.DB
}

func NewMicroserviceRepository(db *gorm.DB) MicroserviceRepositoryInterface {
	return &microserviceRepository{db}
}

func (repository *microserviceRepository) InsertMicroservice(microservice *entities.Microservice) (*entities.Microservice, error) {
	err := repository.db.Table("microservice").Create(microservice).Error

	if err != nil {
		return nil, err
	}

	return microservice, nil
}
