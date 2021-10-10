package test

import (
	"fmt"
	"testing"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain"
	"github.com/bxcodec/faker/v3"
)

func TestNewMicroservice(t *testing.T) {
	m_service := domain.NewMicroservice(faker.Name(), "http", "172.0.0.1", "demo", "demo")
	fmt.Println(m_service)
}
