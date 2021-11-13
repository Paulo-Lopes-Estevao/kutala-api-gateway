package injection

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
)

func (i *injection) NewUserController() controller.UserControllerInterface {
	return controller.NewUserController(i.NewUserUseCase())
}

func (i *injection) NewUserUseCase() usecase.UserUseCaseInterface {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *injection) NewUserRepository() repository.UserRepositoryInterface {
	return repository.NewUserRepository(i.db)
}
