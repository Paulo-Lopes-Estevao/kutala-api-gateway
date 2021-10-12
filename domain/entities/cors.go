package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Cors struct {
	Id              string       `json:"id"`
	Alloworigins    string       `json:"alloworigins"`
	Id_microservice Microservice `json:"id_microservice"`
	State           bool         `json:"state"`
	Created         time.Time    `json:"created"`
	Updated         time.Time    `json:"updated"`
}

func NewCors(alloworigins string, id_microservice Microservice) (*Cors, error) {

	cors := &Cors{
		Alloworigins:    alloworigins,
		Id_microservice: id_microservice,
	}

	err := cors.validate()

	if err != nil {
		return nil, err
	}

	return cors, nil

}

func (cors *Cors) validate() error {

	_, err := govalidator.ValidateStruct(cors)

	if err != nil {
		return err
	}

	return nil
}
