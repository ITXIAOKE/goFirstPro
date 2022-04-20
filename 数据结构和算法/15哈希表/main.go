package main

import (
	"fmt"
	"os"
)

/**
google面试题：

当有一个新员工，要求将该员工的信息加入（id,性别，年龄，住址.....），当输入该员工的id时候，要求找到该员工的信息

要求：
1，不使用数据库，减少内存，速度越快越好，===》哈希表（散列）
2，id从低到高

使用hashtable存放多条部门链表emplink（一个公司有多个部门）
一个部门链表emplink存放多个员工emp（每个部门有多个员工）
一个emp有员工的struct信息（每个员工有自己的信息）

*/
//定义emp

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

//方法待定

func (this *Emp) ShowMe() {
	fmt.Printf("雇员所在链表：%d  找到该雇员id是：%d ，他的名字是：%s\n", this.Id%7, this.Id, this.Name)
}

//定义EmpLink
//我们这里的EmpLink不带表头，即第一个节点就存放雇员

type EmpLink struct {
	Head *Emp
}

//方法待定
//1，添加员工的方法,保证添加时，编号是从小到大

func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head //这个是辅助指针
	var pre *Emp     //辅助指针 pre始终在cur前面

	//如果当前的Emplink就是一个空链表
	if cur == nil {
		this.Head = emp
		return
	}
	//不是空链表，给emp找到对应的位置并插入
	//思路是 让cur和emp比较，然后让pre保持在cur前面
	for {
		if cur != nil {
			//比较
			if cur.Id > emp.Id {
				//找到位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	//退出时，我们将看下是否将emp添加到链表最后
	if pre != nil {
		pre.Next = emp
		emp.Next = cur
	}
}

//显示链表的信息
func (this *EmpLink) showLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空 \n", no)
		return
	}

	//遍历当前的链表，并显示数据
	cur := this.Head //辅助的指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d  名字=%s -->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() //换行处理
}

//根据id查找对应的雇员，如果没有就返回nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//定义hashtable,含有一个链表数组

type HashTable struct {
	LinkArr [7]EmpLink //存储7个链表（部门）,可以理解为一个公司有好多部门，添加进来的员工要进入到具体的部门
}

//给hashtable编写insert 雇员的方法

func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

//显示hashtable所有的雇员
func (this *HashTable) showAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].showLink(i)
	}
}

//编写一个散列方法

func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是对应的链表的下标
}

//编写一个方法，完成查找

func (this *HashTable) FindById(id int) *Emp {
	//使用散列函数，确定将该雇员应该在哪个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

func main() {
	var key, name string
	var id int
	var hashTable HashTable
	for {
		fmt.Println("==========雇员系统菜单=================")
		fmt.Println("intput 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 表示退出雇员")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.showAll()
		case "find":
			fmt.Println("请出入id号：")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				//显示雇员信息
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}

	}
}
