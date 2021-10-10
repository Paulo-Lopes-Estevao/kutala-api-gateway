package test

import (
	"fmt"
	"testing"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/entities"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/database"
	"github.com/bxcodec/faker/v3"
)

var m_service = entities.NewMicroservice(faker.Name(), "demo", "demo")

var conectdb = database.ConnectionDB()
var repositories = repository.MicroserviceRepository{Db: conectdb}
var usecases = usecase.MicroserviceUseCase{MicroserviceRepository: repositories}

func TestNewMicroservice(t *testing.T) {
	fmt.Println(m_service)
}

func TestRepositoryInsertMicroservice(t *testing.T) {
	defer conectdb.Close()
	repositories.InsertMicroservice(m_service)

}

func TestUsecaseCreateMicroservice(t *testing.T) {
	defer conectdb.Close()
	result, _ := usecases.CreateMicroService(m_service)
	fmt.Println(result)

}
