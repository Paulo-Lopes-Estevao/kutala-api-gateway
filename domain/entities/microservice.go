package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Microservice struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Api     string    `json:"api"`
	Path    string    `json:"path"`
	Method  string    `json:"method"`
	Iduser  string    `json:"iduser"`
	User    *User     `json:"user"`
	State   bool      `json:"state"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func NewMicroservice(name string, version string, api string, path string, method string, iduser string) (*Microservice, error) {

	microservice := &Microservice{
		Name:    name,
		Version: version,
		Api:     api,
		Path:    path,
		Method:  method,
		Iduser:  iduser,
		State:   true,
		Created: time.Now(),
		Updated: time.Now(),
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
