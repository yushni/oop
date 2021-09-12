package main

import "fmt"

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

type dog struct {
	breed string
}

func (d dog) sayHello() {
	fmt.Printf("Всім гав гав, я пес! Моя порода %s\n", d.breed)
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
	yurii := student{name: "Юра", age: 24}
	ivan := student{name: "Іван", age: 19}
	petro := student{name: "Петро", age: 32}

	//rex := dog{breed: "вівчарка"}

	startAdultParty(ivan, yurii, petro)
}
