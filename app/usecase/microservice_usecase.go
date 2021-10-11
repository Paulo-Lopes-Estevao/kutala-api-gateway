package usecase

import (
	"fmt"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type MicroserviceUseCaseInterface interface {
	CreateMicroService(microservice *entities.Microservice) (*entities.Microservice, error)
	Auth(username string, password string, microservice *entities.Microservice) (*entities.Microservice, error)
}

type microserviceUseCase struct {
	microserviceRepository repository.MicroserviceRepositoryInterface
}

func NewMicroserviceUseCase(repository repository.MicroserviceRepositoryInterface) MicroserviceUseCaseInterface {
	return &microserviceUseCase{repository}
}

func (usecase *microserviceUseCase) CreateMicroService(microservice *entities.Microservice) (*entities.Microservice, error) {

	data, _ := entities.NewMicroservice(microservice.Name, microservice.Username, microservice.Password)

	microservices, err := usecase.microserviceRepository.InsertMicroservice(data)

	if err != nil {
		return nil, err
	}

	return microservices, nil
}

func (usecase *microserviceUseCase) Auth(username string, password string, microservice *entities.Microservice) (*entities.Microservice, error) {
	data, err := usecase.microserviceRepository.FindUsernameMicroservice(username, microservice)

	if err != nil {
		return nil, fmt.Errorf("The password is invalid for the username")
	}

	if data.VerifyPassword(password) {
		return data, nil
	}

	return nil, fmt.Errorf("The password is invalid for the username")
}
