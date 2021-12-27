package repository

import (
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/usecase/repository"
	"github.com/jinzhu/gorm"
)

type microserviceRepository struct {
	db *gorm.DB
}

func NewMicroserviceRepository(db *gorm.DB) repository.MicroserviceRepositoryInterface {
	return &microserviceRepository{db}
}

func (repository *microserviceRepository) InsertMicroservice(microservice *entities.Microservice) (*entities.Microservice, error) {
	err := repository.db.Table("microservices").Create(microservice).Error

	if err != nil {
		return nil, err
	}

	return microservice, nil
}

func (repository *microserviceRepository) FindUsernameMicroservice(username string, microservice *entities.Microservice) (*entities.Microservice, error) {
	err := repository.db.Table("microservices").Find(microservice, "username = ?", username).Error

	if err != nil {
		return nil, err
	}

	return microservice, nil

}

func (repository *microserviceRepository) FindPathMicroservice(path string, microservice *entities.Microservice) (*entities.Microservice, error) {
	err := repository.db.Table("microservices").Find(microservice, "name = ?", path).Error

	if err != nil {
		return nil, err
	}

	return microservice, nil

}
