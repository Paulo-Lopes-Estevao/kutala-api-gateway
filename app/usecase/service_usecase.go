package usecase

import (
	"fmt"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type ServiceUseCaseInterface interface {
	CreateService(Service *entities.Service) (*entities.Service, error)
	Auth(username string, password string, Service *entities.Service) (*entities.Service, error)
	SearchUuid(uuid string, Service *entities.Service) (*entities.Service, error)
}

type serviceUseCase struct {
	serviceRepository repository.ServiceRepositoryInterface
}

func NewServiceUseCase(repository repository.ServiceRepositoryInterface) ServiceUseCaseInterface {
	return &serviceUseCase{repository}
}

func (usecase *serviceUseCase) CreateService(service *entities.Service) (*entities.Service, error) {

	data, err := entities.NewService(service.Name, service.Username, service.Password)

	if err != nil {
		return nil, err
	}

	services, err := usecase.serviceRepository.InsertService(data)

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (usecase *serviceUseCase) Auth(username string, password string, service *entities.Service) (*entities.Service, error) {
	data, err := usecase.serviceRepository.FindUsernameService(username, service)

	if err != nil {
		return nil, fmt.Errorf("The password is invalid for the username")
	}

	if data.VerifyPassword(password) {
		return data, nil
	}

	return nil, fmt.Errorf("The password is invalid for the username")
}

func (usecase *serviceUseCase) SearchUuid(uuid string, service *entities.Service) (*entities.Service, error) {
	data, err := usecase.serviceRepository.FindIdService(uuid, service)

	if err != nil {
		return nil, fmt.Errorf("Id is invalid")
	}

	return data, nil
}
