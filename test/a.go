package main

import (
	"gopkg.in/gomail.v2"
)

type studentNormal struct {
	name string
	age  int
}

type studentPointer struct {
	name *string
	age  *int
}

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "19725912@qq.com")
	m.SetHeader("To", "19725912@qq.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "dfsdf3f3f3fwef")

	d := gomail.NewDialer("smtp.qq.com",
		465, "19725912@qq.com", "ejusnukrlniabgdd")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

//88:UpfPUMJGrbBWx9VR
//qq:ejusnukrlniabgdd
