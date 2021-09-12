package not_encapsulated

import "fmt"

type Student struct {
	NAME string
	AGE  int
}

func NewStudent(name string, age int) Student {
	return Student{NAME: name, AGE: age}
}

func (s Student) IsAdult() bool {
	return s.AGE > 21
}

func (s Student) SayHello() {
	fmt.Printf("Всім привіт, я спудей! Мене звати %s, мені %d років\n", s.NAME, s.AGE)
}
