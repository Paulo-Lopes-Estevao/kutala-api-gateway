package repository

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain"
	"github.com/jinzhu/gorm"
)

type MicroserviceRepositoryInterface interface {
	InsertMicroservice(microservice domain.Microservice) (*domain.Microservice, error)
	FindMicroservice(name string, username string, password string) (bool, error)
}

type MicroserviceRepository struct {
	Db *gorm.DB
}

func (repository MicroserviceRepository) InsertMicroservice(microservice *domain.Microservice) (*domain.Microservice, error) {
	err := repository.Db.Table("microservice").Create(microservice).Error

	if err != nil {
		return nil, err
	}

	return microservice, nil
}
