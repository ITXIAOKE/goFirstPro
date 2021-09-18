package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int //队首
	tail    int //队尾
}

//入队列
func (c *CircleQueue) Push(val int) (err error) {
	if c.IsFull() {
		return errors.New("queue full")
	}
	//分析出c.tail，在队列尾部，不包含最后的元素
	c.array[c.tail] = val //把值给尾部
	//c.tail++//报错，解决方法如下
	c.tail = (c.tail + 1) % c.maxSize
	return
}

//出队列
func (c *CircleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	//取出，head是指向队首，并且含队首元素
	val = c.array[c.head]
	//c.head++//报错,解决方法如下
	c.head = (c.head + 1) % c.maxSize
	return
}

//显示环形队列
func (c *CircleQueue) ListQueue() {
	size := c.Size() //取出当前队列有多少哥元素show
	if size == 0 {
		fmt.Println("队列为空")
	}
	tempHead := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, c.array[tempHead])
		tempHead = (tempHead + 1) % c.maxSize
	}
	fmt.Println()
}

//判断环形队列为满
func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

//判断环形队列为空
func (c *CircleQueue) IsEmpty() bool {
	return c.tail == c.head
}

//取出环形队列有多少哥元素
func (c *CircleQueue) Size() int {
	return (c.tail + c.maxSize - c.head) % c.maxSize //这个很重要
}

func main() {
	/**
	环形队列
	*/
	//先创建一个队列实例，并初始化数据
	myCirclequeue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
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
			err := myCirclequeue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok!")
			}
		case "get":
			val, err := myCirclequeue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("队列取的数据==", val)
			}
		case "show":
			myCirclequeue.ListQueue()
		case "exit":
			os.Exit(0)
		}

	}

}
