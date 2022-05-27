package main

func main() {
	var a *string
	var b string
	*a = "111"
	b = "222"
	if *a != b {
		return
	}
}
