package entities

import (
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type Microservice struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Version  string    `json:"version"`
	Api      string    `json:"api"`
	Endpoint string    `json:"endpoint"`
	Type     string    `json:"type"`
	Method   string    `json:"method"`
	Iduser   string    `json:"iduser"`
	User     *User     `json:"user"`
	State    bool      `json:"state"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func NewMicroservice(name string, version string, api string, endpoint string, typer string, method string, iduser string) (*Microservice, error) {

	typerUpper := upperText(typer)

	microservice := &Microservice{
		Name:     name,
		Version:  version,
		Api:      api,
		Type:     typerUpper,
		Endpoint: endpoint,
		Method:   method,
		Iduser:   iduser,
		State:    true,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	err := microservice.validate()

	if err != nil {
		return nil, err
	}

	return microservice, nil

}

func upperText(text string) string {
	return strings.ToUpper(text)
}

func (microservice *Microservice) VerifyTypeProtocolComunication(protocol string) bool {
	typer := []string{"GRAPHQL", "REST"}

	protocolUpper := upperText(protocol)

	for _, v := range typer {
		if protocolUpper == v {
			return true
		}
	}

	return false

}

func (microservice *Microservice) validate() error {

	_, err := govalidator.ValidateStruct(microservice)

	if err != nil {
		return err
	}

	return nil
}
