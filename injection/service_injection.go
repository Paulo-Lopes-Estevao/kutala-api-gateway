package injection

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
)

func (i *injection) NewServiceController() controller.ServiceControllerInterface {
	return controller.NewServiceController(i.NewServiceUseCase())
}

func (i *injection) NewServiceUseCase() usecase.ServiceUseCaseInterface {
	return usecase.NewServiceUseCase(i.NewServiceRepository())
}

func (i *injection) NewServiceRepository() repository.ServiceRepositoryInterface {
	return repository.NewServiceRepository(i.db)
}
