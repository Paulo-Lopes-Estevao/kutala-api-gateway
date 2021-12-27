package injection

import (
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/interface/controller"
	"github.com/jinzhu/gorm"
)

type injection struct {
	db *gorm.DB
}

type InjectionInterface interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) InjectionInterface {
	return &injection{db}
}

func (i *injection) NewAppController() controller.AppController {
	return controller.AppController{
		Microservice: i.NewMicroserviceController(),
		User:         i.NewUserController(),
	}
}
