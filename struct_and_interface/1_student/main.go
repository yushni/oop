package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func (s student) sayHello() {
	fmt.Printf("Всім привіт, я студент! Мене звати %s, мені %d років\n", s.name, s.age)
}

func (s student) isAdult() bool {
	return s.age > 21
}

func main() {
	yurii := student{
		name: "Юра",
		age:  24,
	}

	yurii.sayHello()
}
