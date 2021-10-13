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

	data, err := entities.NewMicroservice(microservice.Api, microservice.Method, microservice.Id_service)

	if err != nil {
		return nil, err
	}

	microservices, err := usecase.microserviceRepository.InsertMicroservice(data)

	if err != nil {
		return nil, err
	}

	return microservices, nil
}
