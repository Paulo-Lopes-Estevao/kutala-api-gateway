package usecase

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain"
)

type MicroserviceUseCase struct {
	MicroserviceRepository repository.MicroserviceRepository
}

func (usecase MicroserviceUseCase) CreateMicroService(microservice *domain.Microservice) (*domain.Microservice, error) {
	microservices, err := usecase.MicroserviceRepository.InsertMicroservice(microservice)

	if err != nil {
		return nil, err
	}

	return microservices, nil
}
