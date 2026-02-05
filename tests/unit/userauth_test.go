package adapters_test

import (
	"app/internal/infrastructure/adapters"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := adapters.Static.GenerateJWT("h00", "user")
	if err != nil {
		t.Fatalf("Erro ao gerar token: %v", err)
	}
	valid, claims := adapters.Static.ValidateJWT(token)
	if !valid {
		t.Error("Token inválido")
	}
	if claims.UserID != "h00" || claims.Role != "user" {
		t.Error("Claims inválidos")
	}
	t.Log("Tudo certo!!")
}
