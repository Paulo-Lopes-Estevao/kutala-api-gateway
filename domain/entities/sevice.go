package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Service struct {
	Id              string `json:"id"`
	Scheme          string `json:"scheme"`
	Host            string `json:"host"`
	Path            string `json:"path"`
	Header          string `json:"header"`
	Method          string `json:"method"`
	Id_microservice Microservice
	State           bool      `json:"state"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
}

func NewService(scheme string, host string, path string, header string, method string, id_microservice Microservice) (*Service, error) {

	service := &Service{
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

	err := service.validate()

	if err != nil {
		return nil, err
	}

	return service, nil

}

func (service *Service) validate() error {

	_, err := govalidator.ValidateStruct(service)

	if err != nil {
		return err
	}

	return nil
}
