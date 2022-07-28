package main

import "gopkg.in/gomail.v2"

func main() {
	e := "19725912@qq.com"
	sub := "123"
	body := "dfkjk"

	test(e, sub, body)
}

func test(e string, sub string, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", e)
	m.SetHeader("To", e)
	m.SetHeader("Subject", sub)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.qq.com", 465, "19725912@qq.com", "ejusnukrlniabgdd")
	err := d.DialAndSend(m)
	if err != nil {

	}
}

//88:UpfPUMJGrbBWx9VR
//qq:ejusnukrlniabgdd
