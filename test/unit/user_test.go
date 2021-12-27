package entities_test

import (
	"testing"

	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/domain/entities"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	t.Parallel()
	_, err := entities.NewUser(faker.Name(), "demo", "demo")
	require.Error(t, err)

}
