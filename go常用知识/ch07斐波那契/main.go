package main

import "fmt"

//闭包和递归和动态规划三种方式实现
func main() {
	f := fanacc()
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}

	fmt.Println(fib(5))
	fmt.Println(climbStairs(5))
}

//闭包的方式
func fanacc() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//递归的方式
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

//动态规划来做
func climbStairs(n int)int{
	if n<0{
		return 0
	}
	if n==1{
		return 1
	}
	res:=make([]int,n+1)
	res[0]=0
	res[1]=1
	res[2]=2
	for i := 2; i <=n ; i++ {
		res[i]=res[i-1]+res[i-2]
	}
	return res[n]
}
