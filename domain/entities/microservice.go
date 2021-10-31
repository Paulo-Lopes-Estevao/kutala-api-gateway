package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Microservice struct {
	Id        string    `json:"id"`
	Api       string    `json:"api"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	Idservice string    `json:"idservice"`
	Service   []Service `json:"service"`
	State     bool      `json:"state"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

func NewMicroservice(api string, path string, method string, idservice string) (*Microservice, error) {

	microservice := &Microservice{
		Api:       api,
		Path:      path,
		Method:    method,
		Idservice: idservice,
		State:     true,
		Created:   time.Now(),
		Updated:   time.Now(),
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
