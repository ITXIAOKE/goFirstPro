package main

import "fmt"

//链表实现栈，并使用 interface{} 来代替 int实现多种类型的兼容
type node struct {
	Val  interface{}
	Prev *node //前一个元素
}

type stackLink struct {
	Top    *node //栈顶
	Length int
}

func (sl *stackLink) push(value interface{}) {
	newNode := &node{
		Val:  value,
		Prev: sl.Top}
	sl.Top = newNode
	sl.Length++
}

func (sl *stackLink) pop() interface{} {
	topNodeVal := sl.Top.Val
	sl.Top = sl.Top.Prev
	sl.Length--
	return topNodeVal
}

func (sl stackLink) length() int {
	return sl.Length
}

func (sl stackLink) isEmpty() bool {
	return sl.Length == 0
}

func (sl stackLink) peer() interface{} {
	return sl.Top.Val
}

func stackCompute(str string) int {
	var (
		vsk  stackLink //值栈
		opsk stackLink //操作栈
	)
	for _, v := range str {
		if v <= '9' && v >= '0' {
			vsk.push(int(v) - '0') //值栈中存储字符对应的十进制数
		} else if v == '+' || v == '-' || v == '*' || v == '/' {
			opsk.push(int(v)) //操作栈中存储字符对应的十进制数
		} else if v == ')' {
			n1 := vsk.pop().(int)
			n2 := vsk.pop().(int)
			op := opsk.pop().(int)
			var ans int
			switch op {
			case '+':
				ans = n2 + n1
			case '-':
				ans = n2 - n1
			case '*':
				ans = n2 * n1
			case '/':
				ans = n2 / n1
			}
			vsk.push(int(ans))
		}
	}
	for !opsk.isEmpty() {
		n1 := vsk.pop().(int)
		n2 := vsk.pop().(int)
		op := opsk.pop().(int)
		var ans int
		switch op {
		case '+':
			ans = n2 + n1
		case '-':
			ans = n2 - n1
		case '*':
			ans = n2 * n1
		case '/':
			ans = n2 / n1
		}
		vsk.push(int(ans))
	}
	char := vsk.pop().(int)
	return int(char)
}

/**
我们维护两个栈，一个是值栈，一个是操作栈，我们在读取表达式的时候采取如下的策略：

1，如果遇到 '('，我们将忽略它
2，如果遇到数字，将其压入值栈
3，如果遇到操作符，将其压入操作栈
4，如果遇到 ')'，我们从值栈中取出两个值 n1和 n2，在操作栈中，我们取出一个操作符 op
5，我们进行 n2 op n1的操作（例如 n1 = 3，n2 = 9，op = '/'，我们将执行 9/3 ）
6，将所得的结果压入值栈

*/

func main() {
	str := "(1+((2+3)*(4*5))-6/7)"
	fmt.Println(stackCompute(str))
	fmt.Println("===============")

	fmt.Println(int(3))
	fmt.Println(int(3) - 0)
	fmt.Println(int('3') - 0)   //'3'这个字符对应的十进制是51
	fmt.Println(int('3') - '0') //'0'这个字符对应的十进制是48
}
