package repository

import "github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"

type MicroserviceRepositoryInterface interface {
	InsertMicroservice(microservice *entities.Microservice) (*entities.Microservice, error)
	FindUsernameMicroservice(username string, microservice *entities.Microservice) (*entities.Microservice, error)
	FindPathMicroservice(id string, microservice *entities.Microservice) (*entities.Microservice, error)
}
