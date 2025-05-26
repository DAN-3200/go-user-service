package model_test

import (
	"app/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// assert : averigua e deixa continuar
// require : averigua e não deixa continuar

func Test_UserValidate(t *testing.T) {
	tempoAtual := time.Now().Format("0000-00-00 00:00:00")
	user := model.NewUser(
		"person",
		"person@gmail.com",
		"person321",
		"user",
		tempoAtual,
	)

	err := user.Validate()
	require.NoError(t, err, err)
}

func Test_LoginValidate(t *testing.T) {
	login := model.LoginFields{
		"person@hotmail.com",
		"person321",
	}

	err := login.Validate()
	require.NoError(t, err, err)
}

func Test_LoginValidateFields(t *testing.T) {
	login := model.LoginFields{
		"person@hotmail.com",
		"person",
	}

	errList, err := login.ValidateFields()
	require.NoError(t, err, errList)
}
