package model_test

import (
	"app/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// assert : averigua e deixa continuar
// require : averigua e não deixa continuar

func TestModel(t *testing.T) {
	tempoAtual := time.Now().Format("0000-00-00 00:00:00")
	user := model.NewUser(
		"person",
		"person@gmail.com",
		"person321",
		"user",
		tempoAtual,
	)

	expect := &model.User{
		0,
		"person",
		"person@gmail.com",
		"person321",
		"user",
		tempoAtual,
	}

	require.Equal(t, expect, user)
}
