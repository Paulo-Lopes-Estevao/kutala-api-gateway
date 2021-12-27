package injection

import (
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/interface/controller"
	interfaceRepository "github.com/Paulo-Lopes-Estevao/kutala-api-gateway/interface/repository"
	usecaseRepository "github.com/Paulo-Lopes-Estevao/kutala-api-gateway/usecase/repository"
)

func (i *injection) NewMicroserviceController() controller.MicroserviceControllerInterface {
	return controller.NewMicroserviceController(i.NewUserUseCase())
}

func (i *injection) NewMicroserviceRepository() usecaseRepository.MicroserviceRepositoryInterface {
	return interfaceRepository.NewMicroserviceRepository(i.db)
}
