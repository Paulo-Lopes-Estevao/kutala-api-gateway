package injection

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	interfaceRepository "github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/usecase/interactor"
	usecaseRepository "github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/usecase/repository"
)

func (i *injection) NewUserController() controller.UserControllerInterface {
	return controller.NewUserController(i.NewUserUseCase())
}

func (i *injection) NewUserUseCase() interactor.UserUseCaseInterface {
	return interactor.NewUserUseCase(i.NewUserRepository(), i.NewMicroserviceRepository())
}

func (i *injection) NewUserRepository() usecaseRepository.UserRepositoryInterface {
	return interfaceRepository.NewUserRepository(i.db)
}
