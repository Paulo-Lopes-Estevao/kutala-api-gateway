package adapter

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/database"
)

type microserviceAdapterUsecase struct {
	MicroserviceUseCase usecase.MicroserviceUseCase
}

func ServiceAdapter() *microserviceAdapterUsecase {
	return &microserviceAdapterUsecase{}
}

func MicroserviceAdapter() {
	db := database.ConnectionDB()
	ServiceAdapter := ServiceAdapter()
	ServiceAdapterRepository := repository.MicroserviceRepository{Db: db}
	ServiceAdapter.MicroserviceUseCase = usecase.MicroserviceUseCase{MicroserviceRepository: ServiceAdapterRepository}
}
