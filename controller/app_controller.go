package controller

type AppController struct {
	Microservice interface {
		MicroserviceControllerInterface
	}
}
