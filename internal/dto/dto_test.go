package dto_test

import (
	"app/internal/dto"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UserValidate(t *testing.T) {
	user := dto.UserReq{
		Name:     "person",
		Email:    "person@gmail.com",
		Password: "person321",
		Role:     "user",
	}

	err := user.ValidateFields()
	require.NoError(t, err, err)
}

func Test_LoginValidateFields(t *testing.T) {
	login := dto.Login{
		"person@hotmail.com",
		"person",
	}

	err := login.ValidateFields()
	require.NoError(t, err, err)
}
