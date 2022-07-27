package main

import "fmt"

type studentNormal struct {
	name string
	age  int
}

type studentPointer struct {
	name *string
	age  *int
}

var oldStudents = []*studentNormal{
	{name: "tom", age: 20},
	{name: "sam", age: 22},
	{name: "bill", age: 25}}

var newStudents []studentPointer

func main() {
	for _, v := range oldStudents {
		if v.age > 20 {
			fmt.Println(&v.name)
			var x studentPointer
			x.name = &v.name
			x.age = &v.age
			newStudents = append(newStudents, x)
		}
	}
	fmt.Println(*newStudents[0].name, *newStudents[1].name)
}
