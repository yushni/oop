package main

import "bufio"

// check https://golang.org/doc/effective_go#embedding for more explanations

type reader interface {
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
	f := fileReadWriter{reader: &bufio.Reader{}, writer: &bufio.Writer{}}
	if _, err := f.Read([]byte{}); err != nil {
		panic(err)
	}

	if _, err := f.Write([]byte{}); err != nil {
		panic(err)
	}
}
