package usecase

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type MicroserviceUseCase struct {
	MicroserviceRepository repository.MicroserviceRepository
}

func (usecase MicroserviceUseCase) CreateMicroService(microservice *entities.Microservice) (*entities.Microservice, error) {
	microservices, err := usecase.MicroserviceRepository.InsertMicroservice(microservice)

	if err != nil {
		return nil, err
	}

	return microservices, nil
}
