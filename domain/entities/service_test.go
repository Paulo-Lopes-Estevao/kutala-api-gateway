package entities_test

import (
	"testing"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestNewService(t *testing.T) {
	t.Parallel()
	_, err := entities.NewService(faker.Name(), "demo", "demo")
	require.Error(t, err)

}
