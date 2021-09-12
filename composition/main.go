package main

import (
	"io"
	"os"
)

type reader interface {
	Read(p []byte) (n int, err error)
}

type writer interface {
	Write(p []byte) (n int, err error)
}

type fileCopier struct {
	r reader
	w writer
}

func (c *fileCopier) copy() error {
	content, err := io.ReadAll(c.r)
	if err != nil {
		return err
	}

	if _, err := c.w.Write(content); err != nil {
		return err
	}
	return nil
}

func main() {
	from, err := os.Open("composition/test.txt")
	if err != nil {
		panic(err)
	}
	to, err := os.Create("composition/response.txt")
	if err != nil {
		panic(err)
	}

	fc := fileCopier{r: from, w: to}
	if err := fc.copy(); err != nil {
		panic(err)
	}
}
