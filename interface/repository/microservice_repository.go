package repository

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/usecase/repository"
	"github.com/jinzhu/gorm"
)

type microserviceRepository struct {
	db *gorm.DB
}

func NewMicroserviceRepository(db *gorm.DB) repository.MicroserviceRepositoryInterface {
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

func (repository *microserviceRepository) FindPathMicroservice(path string, microservice *entities.Microservice) (*entities.Microservice, error) {
	err := repository.db.Table("microservice").Find(microservice, "path = ?", path).Error

	if err != nil {
		return nil, err
	}

	return microservice, nil

}
