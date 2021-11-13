package controller

type AppController struct {
	Microservice interface {
		MicroserviceControllerInterface
	}

	User interface {
		UserControllerInterface
	}
}
