package util

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "19725912@qq.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.qq.com", 465, "19725912@qq.com", "ejusnukrlniabgdd")
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
