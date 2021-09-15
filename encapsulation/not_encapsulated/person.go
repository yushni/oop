package not_encapsulated

import "fmt"

type Student struct {
	Name string
	Age  int
}

func NewStudent(name string, age int) Student {
	return Student{Name: name, Age: age}
}

func (s Student) IsAdult() bool {
	return s.Age > 21
}

func (s Student) SayHello() {
	fmt.Printf("Всім привіт, я спудей! Мене звати %s, мені %d років\n", s.Name, s.Age)
}
