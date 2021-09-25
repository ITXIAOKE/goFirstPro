package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s: message %d", name, i)
			i++
		}
	}()
	return c
}

func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch) //注意这里一定要传这个ch，否则只会显示最后的那个值,这里是一个坑
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

func main() {
	m1 := msgGen("chan aa")
	m2 := msgGen("chan bb")
	m3 := msgGen("chan cc")
	m := fanIn(m1, m2, m3)
	for {
		fmt.Println(<-m)
	}
}
