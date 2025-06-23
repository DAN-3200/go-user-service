package userauth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKEY = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(userID, userRole string) (string, error) {
	// formato do JWT : Header.Payload.Signature
	tokenString, err := jwt.
		NewWithClaims(
			jwt.SigningMethodHS256, // = Header
			jwt.MapClaims{ // = Payload
				"userID": userID,
				"iss":    "rest.user.service",
				"role":   userRole,
				"exp":    time.Now().Add(time.Hour * 2).Unix(), // expirar: tempo atual + 2 horas
			},
		).
		SignedString(secretKEY) // = Signature

	return tokenString, err
}

type ModelClaims struct {
	UserID string
	Role   string
	Iss    string
}

func ValidateJWT(tokenString string) (bool, ModelClaims) {
	tokenString = RemoveBearerPrefix(tokenString)
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar se o método é do tipo criptográfico HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "Não é do tipo HMAC", nil
		}
		return secretKEY, nil
	})

	if err != nil {
		fmt.Println(err)
		return false, ModelClaims{}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, ModelClaims{
			UserID: claims["userID"].(string),
			Role:   claims["role"].(string),
			Iss:    claims["iss"].(string),
		}
	} else {
		return false, ModelClaims{}
	}
}

func RemoveBearerPrefix(TokenString string) string {
	if after, ok := strings.CutPrefix(TokenString, "Bearer "); ok {
		TokenString = after
	}

	return TokenString
}
