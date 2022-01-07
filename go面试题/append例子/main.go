package main

import "fmt"

//func main(){
//	x:=[]int{1,2,3}
//	fmt.Println(x)
//	y:=x[:2]
//	fmt.Println(x)
//	y=append(y,50)//公用同一块内存，把3变成50
//	fmt.Println(x)
//	y=append(y,60)//发生了扩容，这时候x与y指向不同的内存块，改变y内存块的值，x不变
//	fmt.Println(x)
//	y[0]=100
//	fmt.Println(x)//改变的是y内存块的值，x不变
//
//}

/**
[1 2 3]
[1 2 3]
[1 2 50]
[1 2 50]
[1 2 50]

*/

func main() {
	x := []int{1, 2, 3}

	y := x[:2]
	fmt.Println(y)
	y = append(y, 50)
	fmt.Println(y)
	y = append(y, 60)
	fmt.Println(y)
	y[0] = 100
	fmt.Println(y)

}

/**
[1 2]
[1 2 50]
[1 2 50 60]
[100 2 50 60]

*/

type Obj struct {
	id int
}

func createObjs(n int) []*Obj {
	objs := make([]Obj, n) //一次把连续的内存给分配出来，不用每次都去分配
	refs := make([]*Obj, 0, n)
	for i := 0; i < n; i++ {
		refs = append(refs, &objs[i])
	}
	return refs
}
