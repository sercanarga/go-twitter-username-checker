package helpers

import "os"

func SendErrorComplex(err error) {
	SendSMS(SMS{
		Err: err,
	})

	SendEmail(EmailStruct{
		Host:           os.Getenv("EMAIL_HOST"),
		Port:           os.Getenv("EMAIL_PORT"),
		SenderEmail:    os.Getenv("EMAIL_SENDER"),
		SenderPassword: os.Getenv("EMAIL_PASSWORD"),
		ToEmail:        os.Getenv("EMAIL_TO"),
		Subject:        os.Getenv("EMAIL_ERROR_SUBJECT"),
		Body:           err.Error(),
	})
}

func SendDownAccount(v string) {
	SendSMS(SMS{
		Username: v,
		Message:  "down",
	})

	SendEmail(EmailStruct{
		Host:           os.Getenv("EMAIL_HOST"),
		Port:           os.Getenv("EMAIL_PORT"),
		SenderEmail:    os.Getenv("EMAIL_SENDER"),
		SenderPassword: os.Getenv("EMAIL_PASSWORD"),
		ToEmail:        os.Getenv("EMAIL_TO"),
		Subject:        os.Getenv("EMAIL_SUCCESS_SUBJECT"),
		Body:           v,
	})
}
