package main

import (
	"fmt"

	"epam-mentoring/oop/encapsulation/encapsulated"
	"epam-mentoring/oop/encapsulation/not_encapsulated"
)

func main() {
	e := encapsulated.NewStudent("Іван", 13)
	ne := not_encapsulated.NewStudent("Іван", 13)

	startParty(e)
	startCorruptedParty(ne)

	fmt.Println()
	// показати fmt як працює
}

func startParty(students ...encapsulated.Student) {
	for _, student := range students {
		if student.IsAdult() {
			student.SayHello()
		}
	}
}

func startCorruptedParty(students ...not_encapsulated.Student) {
	for _, student := range students {
		if student.Name == "Іван" {
			student.Age = 132
		}

		if student.IsAdult() {
			student.SayHello()
		}
	}
}
