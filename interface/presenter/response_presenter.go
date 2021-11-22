package presenter

import "github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"

func Response(userservice *entities.Microservice, User *entities.User) *entities.Microservice {

	return &entities.Microservice{
		Id:       userservice.Id,
		Name:     userservice.Name,
		Version:  userservice.Version,
		Api:      userservice.Api,
		Endpoint: userservice.Endpoint,
		Type:     userservice.Type,
		Method:   userservice.Method,
		Iduser:   User.Id,
		User:     User,
	}

}
