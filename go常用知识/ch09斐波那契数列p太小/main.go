package main

import (
	"fmt"
	"io"
	"strings"
)

// 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen struct {
	gen           func() int
	currentReader io.Reader
}

func (g *intGen) Read(p []byte) (n int, err error) {
	err = io.EOF
	if g.currentReader != nil {
		n, err = g.currentReader.Read(p)
	}
	if err == io.EOF {
		next := g.gen()
		s := fmt.Sprintf("%d\n", next)
		g.currentReader = strings.NewReader(s)
		if n == 0 {
			n, err = g.currentReader.Read(p)
		}
	}
	return n, err
}

func main() {
	f := &intGen{
		gen: Fibonacci(),
	}
	for i := 0; i < 10; i++ {
		b := make([]byte, 2)
		n, err := f.Read(b)
		fmt.Printf("%d bytes read: %q. err = %v\n", n, b, err)
	}
}
