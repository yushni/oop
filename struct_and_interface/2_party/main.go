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

type greetings interface {
	sayHello()
}

// todo: спитати шо таке ті "..."
func startParty(greetings ...greetings) {
	for _, greeting := range greetings {
		greeting.sayHello()
	}
}

func main() {
	yurii := student{name: "Юра", age: 24}
	ivan := student{name: "Іван", age: 19}
	petro := student{name: "Петро", age: 32}

	// todo: спитатись, шо буде якшо пес розівчиться казати 'sayHello'
	rex := dog{breed: "вівчарка"}
	startParty(ivan, yurii, petro, rex)
}
