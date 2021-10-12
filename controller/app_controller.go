package controller

type AppController struct {
	Microservice interface {
		MicroserviceControllerInterface
	}

	Service interface {
		ServiceControllerInterface
	}
}
