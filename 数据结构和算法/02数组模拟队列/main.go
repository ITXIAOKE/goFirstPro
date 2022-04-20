package main

import (
	"errors"
	"fmt"
	"os"
)

/**
队列先进先出queue

问题：
1，实现了一个基本的队列结构，但是没有有效的利用数组空间
2，思考如何使用数组完成环形的队列
*/

//使用一个结构体管理队列

type MyQueue struct {
	maxSize int
	array   [5]int //数组---》模拟队列
	front   int    //指向队列的队首
	tail    int    //指向队列的尾部
}

//添加数据到队列

func (q *MyQueue) AddQueue(val int) (err error) {
	if q.tail == q.maxSize-1 {
		return errors.New("queue full")
	}
	q.tail++
	q.array[q.tail] = val
	return
}

//显示队列，找到队首，然后遍历到队尾
func (q *MyQueue) showQueue() {
	fmt.Println("队列当前的情况是：")
	for i := q.front + 1; i <= q.tail; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
	}
	fmt.Println()
}

//从队列中取出数据

func (q *MyQueue) GetQueue() (val int, err error) {
	if q.tail == q.front { //队空
		return -1, errors.New("queue empty")
	}
	q.front++
	val = q.array[q.front]
	return val, err
}

func main() {
	//先创建一个队列实例，并初始化数据
	myqueue := &MyQueue{
		maxSize: 5,
		front:   -1,
		tail:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add   表示添加数据到队列")
		fmt.Println("2.输入get   表示从队列获取数据")
		fmt.Println("3.输入show  表示显示队列")
		fmt.Println("4.输入exit  表示退出队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你想要入队列的数据")
			fmt.Scanln(&val)
			err := myqueue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok!")
			}
		case "get":
			val, err := myqueue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("队列取的数据==", val)
			}
		case "show":
			myqueue.showQueue()
		case "exit":
			os.Exit(0)
		}

	}

}
