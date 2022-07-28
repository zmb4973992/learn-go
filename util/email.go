package util

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "19725912@qq.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.qq.com", 465, "19725912@qq.com", "ejusnukrlniabgdd")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	_ = d.DialAndSend(m)
}
