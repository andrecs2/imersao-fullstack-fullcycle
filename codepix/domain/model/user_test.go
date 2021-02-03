package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/codeedu/imersao/codepix-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPUser(t *testing.T) {
	name := "Joao"
	email := "joao@codepix.com"
	user, err := model.NewUser(name, email)


	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, user.name, name)
	require.Equal(t, user.email, email)

}
