package main

import (
	"fmt"
	"time"
)

/**
思路
回文子序列都是动态规划经典题目,动态规划五部曲：:::
1,确定dp数组以及下标的含义
2,确定递推公式
3,dp数组如何初始化
4,确定遍历顺序
5,列举推导dp数组


确定dp数组以及下标的含义:::
dp[i] [j]:字符串s在[i, j]范围内最长的回文子序列的长度为dp[i] [j].

确定递推公式:::
如果s[i]与s[j]相同，那么dp[i] [j]= dp[i + 1] [j - 1] + 2;
如果s[i]与s[j]不相同，说明s[i]和s[j]的同时加入 并不能增加[i,j]区间回文子串的长度，那么分别加入s[i]、s[j]看看哪一个可以组成最长的回文子序列。
加入s[j]的回文子序列长度为dp[i + 1] [j]。
加入s[i]的回文子序列长度为dp[i] [j - 1]。
那么dp[i][j]一定是取最大的，即：dp[i][j] = max(dp[i + 1] [j], dp[i] [j - 1]);

dp数组如何初始化::::
这里我们使dp[i][j]=1,,也就是当i与j相同，一个字符的回文子序列长度就是1。

确定遍历顺序::::
注意我们二维数组dp[i][j]的定义，如果i>j是没有意义的。
2021-09-14 18-48-20 的屏幕截图
从上面图中可以看出如果我们想求dp[i][j]，那么其他3个必须都是已知的，很明显从上往下遍历是不行的，
我们只能让i从最后一个字符往前遍历，j从i的下一个开始遍历，也就是从下到上，从左到右的顺序，这样才能保证计算d的时候，a，b，c的值都已经计算过了。

列举推导dp数组::::
列举完毕后，发现dp[0][length - 1]即是最后需要的结果返回即可。
*/

func main() {
	fmt.Println(LongPalindrome("babad"))
	fmt.Println(LongPalindrome("dbbd"))
	fmt.Println(LongPalindrome("c"))
	fmt.Println(LongPalindrome("cc"))
	fmt.Println(LongPalindrome("ccc"))
	fmt.Println(LongPalindrome(""))
	fmt.Println(LongPalindrome("aaaaaaa"))
	time.Sleep(200000)

	fmt.Println(LongPalindrome("=================="))

	fmt.Println(longestPalindromeSubseq("babad"))
	fmt.Println(longestPalindromeSubseq("dbbd"))
	fmt.Println(longestPalindromeSubseq("c"))
	fmt.Println(longestPalindromeSubseq("dd"))
	fmt.Println(longestPalindromeSubseq("ccc"))
	fmt.Println(longestPalindromeSubseq(""))
	fmt.Println(longestPalindromeSubseq("aaaaaaa"))
}

//第二中方式：
func longestPalindromeSubseq(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		}
	}

	for i := l - 1; i >= 0; i-- {

		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//第一种方式：：遍历了二维数组，时间复杂度 O(n^2)

func LongPalindrome(s string) (string, int) {

	//回文字符串：bab，，，使用动态规划解题
	n := len(s)
	if n == 0 {
		return "", 0
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
	//fmt.Println(dp)
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

	return ret, maxLen

}
