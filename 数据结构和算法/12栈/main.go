package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int //最大存放数
	Top    int //表示栈顶，栈顶随着入栈的数据而慢慢的增加
	arr    [5]int
}

//入栈

func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full,不能再添加了")
		return errors.New("stack full,不能再添加了")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

//出栈

func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 { //先判断栈是否为空
		fmt.Println("stack empty!!")
		return 0, errors.New("stack empty!!")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

//遍历栈,从栈顶开始遍历，栈是先进后出，后进先出的

func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func main() {
	statck := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	statck.Push(1)
	//statck.Push(2)
	//statck.Push(3)
	//statck.Push(4)
	//statck.Push(5)
	//statck.Push(6)

	statck.List()

	fmt.Println(statck.Pop())
	statck.Pop()
	statck.List()
}
