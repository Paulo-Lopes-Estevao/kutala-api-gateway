package repository

import (
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/usecase/repository"
	"github.com/jinzhu/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepositoryInterface {
	return &serviceRepository{db}
}

func (repository *serviceRepository) InsertUser(User *entities.User) (*entities.User, error) {
	err := repository.db.Table("users").Create(User).Error

	if err != nil {
		return nil, err
	}

	return User, nil
}

func (repository *serviceRepository) FindUsernameService(username string, User *entities.User) (*entities.User, error) {
	err := repository.db.Table("users").Find(User, "username = ?", username).Error

	if err != nil {
		return nil, err
	}

	return User, nil

}

func (repository *serviceRepository) FindIdUser(id string, User *entities.User) (*entities.User, error) {
	err := repository.db.Table("users").Find(User, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return User, nil

}
