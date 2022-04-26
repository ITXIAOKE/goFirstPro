package main

import "fmt"

//字节算法题
//给定一个正整数，任意调整位数顺序，找出比它大且离它最近的数
//比如1234 输出1243
//比如1243 输出1324

//有 1、2、3、4 这四个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？

func main() {
	//threeNum() //1234生成任意三个数的组合
	fourNum() //1234生成任意四个数的组合
}

func fourNum() {
	res := make([]int, 0)
	total := 0
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for k := 1; k < 5; k++ {
				for m := 1; m < 5; m++ {
					if i != j && i != k && i != m && j != k && j != m && m != k {
						total += 1
						res = append(res, i*1000+j*100+k*10+m)
					}
				}
			}
		}
	}

	fmt.Println("一共元素是：", total)
	fmt.Println(res)
	//对res切片排序
	//然后找到1234对应的角标，返回的值是角标+1对应的值

	for i := 0; i < len(res); i++ {
		if res[i] == 1324 {
			fmt.Println("最接进这个值的是：", res[i+1])
		}
	}

}

func threeNum() {
	totalCount := 0
	res := []int{}

	/*以下为三重循环*/
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for k := 1; k < 5; k++ {
				/*确保 i 、j 、k 三位互不相同*/
				if i != k && i != j && j != k {
					totalCount++
					fmt.Println("第", totalCount, "方案", "i =", i, "j =", j, "k =", k)
					res = append(res, i*100+j*10+k)
				}
			}
		}
	}

	fmt.Println("共", totalCount, "种方案")
	fmt.Println("每个元素是：", res)
}
