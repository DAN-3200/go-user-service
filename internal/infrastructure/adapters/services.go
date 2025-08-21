package adapters

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ServiceLayer struct {
	JWT
	SessionCache
}

var Static = ServiceLayer{}

func LayerService() *ServiceLayer {
	return &ServiceLayer{}
}

// [https://www.youtube.com/watch?v=TrdWr3BmqT8]
// Simple Mail Transfer Protocol (smtp)
func (it ServiceLayer) SendMail(to string, body string) error {
	from := os.Getenv("MYEMAIL")
	password := os.Getenv("MYPASSWORD")
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		[]string{to},
		[]byte(body),
	)

	if err != nil {
		return err
	}

	return nil
}

func (it ServiceLayer) GenerateUUID() string {
	return uuid.New().String()
}

func (it ServiceLayer) HashPassword(password string) (string, error) {
	var hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (it ServiceLayer) CompareHashPassword(pivotPassword, inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pivotPassword), []byte(inputPassword))
	if err != nil {
		return fmt.Errorf("Password invalid")
	}
	return nil
}
