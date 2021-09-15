package main

import (
	"fmt"
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

func (c *fileCopier) Read(p []byte) (n int, err error) {
	fmt.Println("то якийсь мій супер класний рідер з складною логікою")
	return c.r.Read(p)
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
	defer from.Close()

	to, err := os.Create("composition/response.txt")
	if err != nil {
		panic(err)
	}
	defer to.Close()

	fc := fileCopier{r: from, w: to}
	if err := fc.copy(); err != nil {
		panic(err)
	}

	// fc.Write([]byte("\n побільше програм"))
}
