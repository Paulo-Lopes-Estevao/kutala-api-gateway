package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type Microservice struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	State    bool      `json:"state"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func NewMicroservice(name string, username string, password string) (*Microservice, error) {
	microservice := &Microservice{
		Name:     name,
		Username: username,
		Password: password,
		State:    true,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	err := microservice.passwordEncrypt()

	if err != nil {
		return nil, err
	}

	return microservice, nil
}

func (microservice *Microservice) passwordEncrypt() error {
	password, err := bcrypt.GenerateFromPassword([]byte(microservice.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	microservice.Password = string(password)

	err = microservice.validate()

	if err != nil {
		return err
	}

	return nil

}

func (microservice *Microservice) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(microservice.Password), []byte(password))
	return err == nil
}

func (microservice *Microservice) validate() error {

	_, err := govalidator.ValidateStruct(microservice)

	if err != nil {
		return err
	}

	return nil
}
