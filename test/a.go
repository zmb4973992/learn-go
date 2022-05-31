package main

import "log"

func main() {
	var a *string
	if a == nil || *a == "hello" {
		log.Print("通过1")
	}
	if *a == "hello" || a == nil {
		log.Print("通过2")
	}

}
