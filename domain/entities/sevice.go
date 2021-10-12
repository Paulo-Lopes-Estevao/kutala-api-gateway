package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	State    bool      `json:"state"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func NewService(name string, username string, password string) (*Service, error) {
	service := &Service{
		Name:     name,
		Username: username,
		Password: password,
		State:    true,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	err := service.passwordEncrypt()

	if err != nil {
		return nil, err
	}

	return service, nil
}

func (service *Service) passwordEncrypt() error {
	password, err := bcrypt.GenerateFromPassword([]byte(service.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	service.Password = string(password)

	err = service.validate()

	if err != nil {
		return err
	}

	return nil

}

func (service *Service) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(service.Password), []byte(password))
	return err == nil
}

func (service *Service) validate() error {

	_, err := govalidator.ValidateStruct(service)

	if err != nil {
		return err
	}

	return nil
}
