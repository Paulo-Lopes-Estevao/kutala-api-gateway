package test

import (
	"fmt"
	"testing"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/database"
	"github.com/bxcodec/faker/v3"
)

func TestNewMicroservice(t *testing.T) {
	m_service := domain.NewMicroservice(faker.Name(), "demo", "demo")
	fmt.Println(m_service)
}

func TestInsertMicroservice(t *testing.T) {
	conectdb := database.ConnectionDB()
	defer conectdb.Close()

	repository := repository.MicroserviceRepository{Db: conectdb}

	m_service := domain.NewMicroservice(faker.Name(), "demo", "demo")

	repository.InsertMicroservice(m_service)

}
