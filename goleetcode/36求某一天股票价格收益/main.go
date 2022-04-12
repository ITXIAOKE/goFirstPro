package main

import "fmt"

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func main() {
	p:=[]int{7,5,3,6,4,9,15}//3+5+6
	fmt.Println(maxProfit(p))
}
