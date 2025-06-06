package userauth_test

import (
	"app/internal/userauth"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := userauth.GenerateJWT("h00", "user")
	if err != nil {
		t.Fatalf("Erro ao gerar token: %v", err)
	}
	valid, claims := userauth.ValidateJWT(token)
	if !valid {
		t.Error("Token inválido")
	}
	if claims.UserID != "h00" || claims.Role != "user" {
		t.Error("Claims inválidos")
	}
	t.Log("Tudo certo!!")
}
