package main

import "fmt"

func main() {
	fmt.Println(LongPalindrome("babad"))
	fmt.Println(LongPalindrome("dbbd"))
	fmt.Println(LongPalindrome("c"))
	fmt.Println(LongPalindrome("cc"))
	fmt.Println(LongPalindrome("ccc"))
	fmt.Println(LongPalindrome(""))
	fmt.Println(LongPalindrome("aaaaaaa"))
}

//遍历了二维数组，时间复杂度 O(n^2)

func LongPalindrome(s string) (string,int) {

	//回文字符串：bab，，，使用动态规划解题
	n := len(s)
	if n == 0 {
		return "",0
	}

	//定义二维数组dp[i][j]用来记录 s[i:j+1]是否是回文字符串
	//用0代表不是，用1代表是
	dp := make([][]int, n)
	//对二维数组初始化
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	//状态转移思路
	//dp[i][j]=1的条件
	//1，j>i
	//2,dp[i+1][j+1]=1
	//3,s[i]==s[j]
	//还有特殊情况：
	//（1）i==j dp[i][j]=1
	//（2）j==i+1 满足s[i]==s[j]则dp[i][j]=1,否则dp[i][j]=

	//dp[i][j]由dp[i+1][j+1]推出，所以i是从大往小遍历，j是从小到大遍历
	//j>=i
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if j == i {
				dp[i][j] = 1
			} else if j == i+1 {
				if s[j] == s[i] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] && dp[i+1][j-1] == 1 {
					dp[i][j] = 1
				}
			}
		}
	}
	fmt.Println(dp)
	//遍历dp数组，找到哪些为1的i,j
	//j-i最大的则是最长的回文子串
	//最长子串
	ret := ""
	//最长长度
	maxLen := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dp[i][j] == 1 && j-i+1 > maxLen {
				//字符串切片不包含end下标，所以j要加一
				ret = s[i : j+1]
				maxLen = j - i + 1
			}
		}
	}

	return ret,maxLen

}
