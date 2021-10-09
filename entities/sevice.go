package entities

import "time"

type Service struct {
	Path            string `json:"path"`
	Id_Microservice Microservice
	Header          string    `json:"header"`
	Method          string    `json:"Method"`
	State           bool      `json:"state"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
}
