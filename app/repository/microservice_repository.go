package repository

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/jinzhu/gorm"
)

type MicroserviceRepositoryInterface interface {
	InsertMicroservice(microservice entities.Microservice) (*entities.Microservice, error)
	FindMicroservice(name string, username string, password string) (bool, error)
}

type MicroserviceRepository struct {
	Db *gorm.DB
}

func (repository MicroserviceRepository) InsertMicroservice(microservice *entities.Microservice) (*entities.Microservice, error) {
	err := repository.Db.Table("microservice").Create(microservice).Error

	if err != nil {
		return nil, err
	}

	return microservice, nil
}
