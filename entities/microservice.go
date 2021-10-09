package entities

import "time"

type Microservice struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Scheme   string    `json:"scheme"`
	Host     string    `json:"host"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	State    bool      `json:"state"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func NewMicroservice(name string, scheme string, host string, username string, password string) *Microservice {
	microservice := &Microservice{
		Name:     name,
		Scheme:   scheme,
		Host:     host,
		Username: username,
		Password: password,
		State:    true,
	}
	return microservice
}
