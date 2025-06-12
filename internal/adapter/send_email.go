package adapter

import (
	"net/smtp"
	"os"
)

type Drive struct{}

func SetDrive() *Drive {
	return &Drive{}
}

// [https://www.youtube.com/watch?v=TrdWr3BmqT8]
// Simple Mail Transfer Protocol (smtp)
func (it Drive) SendMail(to string, body string) error {
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
