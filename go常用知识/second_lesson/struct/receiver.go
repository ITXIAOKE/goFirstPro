package main

import "fmt"

func main() {

	// 因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		Age: 10,
	}
	u.ChangeName("Tom Changed!")
	u.ChangeAge(100)
	fmt.Printf("%v \n", u)

	// 因为 up 指针，所以内部的数据是可以被改变的
	up := &User{
		Name: "xiaoke",
		Age: 12,
	}

	// 因为 ChangeName 的接收器是结构体
	// 所以 up 的数据还是不会变
	up.ChangeName("xiaoke Changed!")
	up.ChangeAge(120)

	fmt.Printf("%v \n", up)
}

type User struct {
	Name string
	Age int
}

// 结构体接收器
func (u User) ChangeName(newName string)  {
	u.Name = newName
}

// 指针接收器
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

/**
go新人3个坑
1：字符串长度计算，使用utf8.RuneCountInString
2: 切片   append后，变成一个新的slice  和数组共用底层数据结构
3: 结构体采用结构体或者指针绑定方法，遇事不决用指针
 */

/**
以下这种用结构体

type Handle func()

func (h Handle)Hello(){

}
 */
