package domain

import "time"

type Service struct {
	Id              string `json:"id"`
	Scheme          string `json:"scheme"`
	Host            string `json:"host"`
	Path            string `json:"path"`
	Header          string `json:"header"`
	Method          string `json:"Method"`
	Id_microservice Microservice
	State           bool   `json:"state"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
}
