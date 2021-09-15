package main

import (
	"bufio"
	"bytes"
)

// check https://golang.org/doc/effective_go#embedding for more explanations
type reader interface {
	// Read todo: спитатисі шо буде якшо повидаляти назви аргументів?
	Read(p []byte) (n int, err error)
}

type writer interface {
	Write(p []byte) (n int, err error)
}

type readWriter interface {
	reader
	writer
}

type fileReadWriter struct {
	reader
	writer
}

func main() {
	// todo: спитати чи можна які імплементації можна використовувати тут?
	f := fileReadWriter{reader: &bufio.Reader{}, writer: &bytes.Buffer{}}
	if _, err := f.Read([]byte{}); err != nil {
		panic(err)
	}

	if _, err := f.Write([]byte{}); err != nil {
		panic(err)
	}
	// todo: які методи ше має 'fileReadWriter'?
	// todo: чи зміниться кількість методів, якшо змінити bytes.Buffer на bufio.Writer{}
}
