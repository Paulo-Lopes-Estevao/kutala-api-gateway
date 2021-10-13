package entities_test

import (

	//"github.com/asaskevich/govalidator"
	"testing"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestNewMicroservice(t *testing.T) {
	t.Parallel()
	_, err := entities.NewMicroservice(faker.URL(), "POST", "111")
	require.Error(t, err)

}
