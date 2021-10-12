package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Microservice struct {
	Id              string    `json:"id"`
	Scheme          string    `json:"scheme"`
	Host            string    `json:"host"`
	Path            string    `json:"path"`
	Header          string    `json:"header"`
	Method          string    `json:"method"`
	Id_microservice Service   `json:"id_microservice"`
	State           bool      `json:"state"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
}

func NewMicroservice(scheme string, host string, path string, header string, method string, id_microservice Service) (*Microservice, error) {

	microservice := &Microservice{
		Scheme:          scheme,
		Host:            host,
		Path:            path,
		Header:          header,
		Method:          method,
		Id_microservice: id_microservice,
		State:           true,
		Created:         time.Now(),
		Updated:         time.Now(),
	}

	err := microservice.validate()

	if err != nil {
		return nil, err
	}

	return microservice, nil

}

func (microservice *Microservice) validate() error {

	_, err := govalidator.ValidateStruct(microservice)

	if err != nil {
		return err
	}

	return nil
}
