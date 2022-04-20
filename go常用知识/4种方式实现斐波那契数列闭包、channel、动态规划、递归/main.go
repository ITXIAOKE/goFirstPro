package main

import "fmt"

//闭包和递归和动态规划三种方式实现
func main() {
	//f := fanacc()
	//for i := 0; i < 5; i++ {
	//	fmt.Println(f())
	//}
	//
	//fmt.Println(fib(5))
	//fmt.Println(climbStairs(5))

	for x := range Fib(10) {
		fmt.Println(x)
	}

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
func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	res[2] = 2
	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}

//采用渠道channel的方式来写

func Fib(n int) chan int { //返回n以内的数列
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}


/**
首先，动态规划问题的一般形式就是求最值。动态规划其实是运筹学的一种最优化方法，只不过在计算机问题上应用比较多，
比如说让你求最长递增子序列呀，最小编辑距离呀等等。

既然是要求最值，核心问题是什么呢？求解动态规划的核心问题是穷举。因为要求最值，肯定要把所有可行的答案穷举出来，然后在其中找最值呗。

动态规划这么简单，就是穷举就完事了？我看到的动态规划问题都很难啊！

首先，动态规划的穷举有点特别，因为这类问题存在「重叠子问题」，如果暴力穷举的话效率会极其低下，
所以需要「备忘录」或者「DP table」来优化穷举过程，避免不必要的计算。

而且，动态规划问题一定会具备「最优子结构」，才能通过子问题的最值得到原问题的最值。

另外，虽然动态规划的核心思想就是穷举求最值，但是问题可以千变万化，穷举所有可行解其实并不是一件容易的事，
只有列出正确的「状态转移方程」，才能正确地穷举。

以上提到的重叠子问题、最优子结构、状态转移方程就是动态规划三要素。
具体什么意思等会会举例详解，但是在实际的算法问题中，写出状态转移方程是最困难的，
这也就是为什么很多朋友觉得动态规划问题困难的原因，我来提供我研究出来的一个思维框架，辅助你思考状态转移方程：

明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 dp 数组/函数的含义。

按上面的套路走，最后的结果就可以套这个框架：

// 伪码
//初始化 base case
dp[0][0][...] = base case
//进行状态转移
for _,状态1 = range 状态1的所有取值{
    for _,状态2 = range 状态2的所有取值{
        for ...{
            dp[状态1][状态2][...] = 求最值(选择1， 选择2， ...)
        }

    }
}
*/
