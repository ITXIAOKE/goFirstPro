package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int //函数

func (g intGen) Read(p []byte) (n int, err error) {//为函数实现接口
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) { //什么样类型才能满足io.reader接口的要求呢？就是实现了Read函数的类型
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() { //可能这里调用Read函数
		fmt.Println(scanner.Text()) //也可能这里调用上面的Read函数
	}
}

func main() {
	var f intGen = Fibonacci()
	printFileContents(f)
}
