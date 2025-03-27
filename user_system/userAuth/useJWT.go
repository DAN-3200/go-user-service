package userAuth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKEY = []byte(os.Getenv("secret_key"))

func GenerateJWT(userID, userEmail, userRole string) (string, error) {

	// formato do JWT : Header.Payload.Signature
	var tokenString, err = jwt.NewWithClaims(
		jwt.SigningMethodHS256, // = Header

		jwt.MapClaims{ // = Payload
			"userID": userID,
			"iss":    userEmail,
			"role":   userRole,
			"exp":    time.Now().Add(time.Hour * 2).Unix(), // expirar: tempo atual + 2 horas
		},
	).SignedString(secretKEY) // = Signature

	return tokenString, err
}

type ModelClaims struct {
	UserID string
	Iss    string
	Role   string
}

func ValidateJWT(tokenString string) (bool, ModelClaims) {
	tokenString = RemoveBearerPrefix(tokenString)

	// Analisa o Token
	var token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar se o método é do tipo criptográfico HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "Não é do tipo HMAC", nil
		}
		// utiliza a `secretKEY` para validação
		return secretKEY, nil
	})

	if err != nil {
		fmt.Println(err)
		return false, ModelClaims{}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, ModelClaims{
			UserID: claims["userID"].(string),
			Iss:    claims["iss"].(string),
			Role:   claims["role"].(string),
		}
	} else {
		return false, ModelClaims{}
	}
}

// Remover o prefixo `Bearer`
func RemoveBearerPrefix(TokenString string) string {
	if strings.HasPrefix(TokenString, "Bearer ") {
		TokenString = strings.TrimPrefix(TokenString, "Bearer ")
	}

	return TokenString
}
