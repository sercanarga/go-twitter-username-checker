package helpers

import "fmt"

type SMS struct {
	Username string
	Message  string
	Err      error
}

func SendSMS(sms SMS) {
	if sms.Err != nil {
		fmt.Println(sms.Err.Error())
		return
	}

	fmt.Println(sms.Username, sms.Message)
}
