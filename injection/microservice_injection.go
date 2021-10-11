package injection

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
)


func(i *injection)NewMicroserviceController()controller.MicroserviceControllerInterface{
	return controller.NewMicroserviceController(i.NewMicroserviceUseCase())
}

func(i *injection)NewMicroserviceUseCase()usecase.MicroserviceUseCaseInterface{
	return usecase.NewMicroserviceUseCase(i.NewMicroserviceRepository())
}

func(i *injection)NewMicroserviceRepository()repository.MicroserviceRepositoryInterface{
	return repository.NewMicroserviceRepository(i.db)
}