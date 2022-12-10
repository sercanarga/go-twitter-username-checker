package helpers

func SendErrorComplex(err error) {
	SendEmail(EmailStruct{
		Host:           "mail.sercanarga.com",
		Port:           465,
		SenderEmail:    "twitter.bot@sercanarga.com",
		SenderPassword: "bg1qs!5bMz65MiClp",
		ToEmail:        "mail@sercanarga.com",
		Subject:        "HATA - Düşen Hesaplar",
		Body:           err.Error(),
	})

	SendSMS(SMS{
		Err: err,
	})
}

func SendDownAccount(v string) {
	SendSMS(SMS{
		Username: v,
		Message:  "düştü",
	})

	SendEmail(EmailStruct{
		Host:           "mail.sercanarga.com",
		Port:           465,
		SenderEmail:    "twitter.bot@sercanarga.com",
		SenderPassword: "bg1qs!5bMz65MiClp",
		ToEmail:        "mail@sercanarga.com",
		Subject:        "Düşen Hesap (" + v + ")",
		Body:           v,
	})
}
