package repository

import "github.com/Paulo-Lopes-Estevao/kutala-api-gateway/domain/entities"

type UserRepositoryInterface interface {
	InsertUser(User *entities.User) (*entities.User, error)
	FindUsernameService(username string, User *entities.User) (*entities.User, error)
	FindIdUser(id string, User *entities.User) (*entities.User, error)
}
