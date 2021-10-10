package domain

import "time"

type Microservice struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	State    bool      `json:"state"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func NewMicroservice(name string, username string, password string) *Microservice {
	microservice := &Microservice{
		Name:     name,
		Username: username,
		Password: password,
		State:    true,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	return microservice
}
