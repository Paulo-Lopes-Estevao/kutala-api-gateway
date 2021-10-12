package repository

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/jinzhu/gorm"
)

type MicroserviceRepositoryInterface interface {
	InsertMicroservice(microservice *entities.Microservice) (*entities.Microservice, error)
	FindUsernameMicroservice(username string, microservice *entities.Microservice) (*entities.Microservice, error)
	FindIdMicroservice(id string, microservice *entities.Microservice) (*entities.Microservice, error)
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

func (repository *microserviceRepository) FindUsernameMicroservice(username string, microservice *entities.Microservice) (*entities.Microservice, error) {
	err := repository.db.Table("microservice").Find(microservice, "username = ?", username).Error

	if err != nil {
		return nil, err
	}

	return microservice, nil

}

func (repository *microserviceRepository) FindIdMicroservice(id string, microservice *entities.Microservice) (*entities.Microservice, error) {
	err := repository.db.Table("microservice").Find(microservice, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return microservice, nil

}
