package encapsulated

import (
	"fmt"
)

type Student struct {
	// спитати який (скоуп) зона видимості полів структури
	name string
	age  int
}

func NewStudent(name string, age int) Student {
	return Student{name: name, age: age}
}

func (s Student) IsAdult() bool {
	return s.age > 21
}

func (s Student) SayHello() {
	fmt.Printf("Всім привіт, я спудей! Мене звати %s, мені %d років\n", s.name, s.age)
}
