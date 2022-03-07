package main

import (
	"fmt"
	"sync"
)

//互斥锁、不可重入锁，在一个goroutine下，重复上锁的话，就会发生panic
var mu sync.Mutex
var chain string

func main() {
	chain = "main"
	A()
	fmt.Println(chain)
}

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + "-->A"
	B()
}

func B() {
	//mu.Lock()
	//defer mu.Unlock()
	chain = chain + "-->B"
	C()
}

func C() {
	//mu.Lock()
	//defer mu.Unlock()
	chain = chain + "-->C"
}

//结果：发生panic

//只一个函数上锁后，结果如下：
//main-->A-->B-->C
