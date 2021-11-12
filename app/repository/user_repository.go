package repository

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/jinzhu/gorm"
)

type UserRepositoryInterface interface {
	InsertUser(User *entities.User) (*entities.User, error)
	FindUsernameService(username string, User *entities.User) (*entities.User, error)
	FindIdUser(id string, User *entities.User) (*entities.User, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &serviceRepository{db}
}

func (repository *serviceRepository) InsertUser(User *entities.User) (*entities.User, error) {
	err := repository.db.Table("User").Create(User).Error

	if err != nil {
		return nil, err
	}

	return User, nil
}

func (repository *serviceRepository) FindUsernameService(username string, User *entities.User) (*entities.User, error) {
	err := repository.db.Table("User").Find(User, "username = ?", username).Error

	if err != nil {
		return nil, err
	}

	return User, nil

}

func (repository *serviceRepository) FindIdUser(id string, User *entities.User) (*entities.User, error) {
	err := repository.db.Table("User").Find(User, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return User, nil

}
