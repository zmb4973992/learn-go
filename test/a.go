package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	var a *string
	var b *string
	str := "hello"
	b = &str
	a, err := Encrypt(b)
	if err != nil {
		return
	}
	log.Print(*a)
}
func Encrypt(originalString *string) (encryptedString *string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*originalString), 10) //10为加密难度，取值范围为4-31，官方建议10
	if err != nil {
		return nil, err
	}
	*encryptedString = string(bytes)
	return encryptedString, nil
}
