package usecase

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type MicroserviceUseCaseInterface interface {
	CreateMicroService(microservice *entities.Microservice) (*entities.Microservice, error)
}

type microserviceUseCase struct {
	microserviceRepository repository.MicroserviceRepositoryInterface
}


func NewMicroserviceUseCase(repository repository.MicroserviceRepositoryInterface) MicroserviceUseCaseInterface {
	return &microserviceUseCase{repository}
}


func (usecase *microserviceUseCase) CreateMicroService(microservice *entities.Microservice) (*entities.Microservice, error) {
	microservices, err := usecase.microserviceRepository.InsertMicroservice(microservice)

	if err != nil {
		return nil, err
	}

	return microservices, nil
}
