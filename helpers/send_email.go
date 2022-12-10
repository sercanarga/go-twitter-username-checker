package helpers

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

type EmailStruct struct {
	Host           string
	Port           string
	SenderEmail    string
	SenderPassword string
	ToEmail        string
	Subject        string
	Body           string
}

func SendEmail(email EmailStruct) (bool, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", email.SenderEmail)
	m.SetHeader("To", email.ToEmail)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text", email.Body)

	port, _ := strconv.Atoi(email.Port)
	d := gomail.NewDialer(email.Host, port, email.SenderEmail, email.SenderPassword)

	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}
