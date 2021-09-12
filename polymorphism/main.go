package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p person) isAdult() bool {
	return p.age > 21
}

type student struct {
	person
}

func (s student) sayHello() {
	fmt.Printf("Всім привіт, я студент! Мене звати %s, мені %d років\n", s.name, s.age)
}

type professor struct {
	person
}

func (p professor) sayHello() {
	fmt.Printf("Всім привіт, я старий старий професор! Мене звати пан %s!\n", p.name)
}

func (p professor) isAdult() bool {
	fmt.Println("А ти шо не бачиш чи шо?")
	return p.person.isAdult()
}

type adultGreetings interface {
	sayHello()
	isAdult() bool
}

func startAdultParty(greetings ...adultGreetings) {
	for _, greeting := range greetings {
		if !greeting.isAdult() {
			continue
		}
		greeting.sayHello()
	}
}

func main() {
	yurii := student{person: person{name: "Юра", age: 24}}
	ivan := student{person: person{name: "Іван", age: 19}}
	petro := professor{person: person{name: "Петро", age: 32}}

	startAdultParty(ivan, yurii, petro)
}
