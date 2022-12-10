package helpers

import "gopkg.in/gomail.v2"

type EmailStruct struct {
	Host           string
	Port           int
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

	d := gomail.NewDialer(email.Host, email.Port, email.SenderEmail, email.SenderPassword)

	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}
