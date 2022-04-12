package main

import "fmt"

func main() {

	var M int //M组任务数据
	var N int //每组机器数量
	var B int //机器配置花费时间
	var J int //机器完成任务花费时间

	takeTimeConfig := make([]int, 0) //花费配置时间
	takeTime := make([]int, 0)       //花费完成任务时间

	for i := 1; i <= M; i++ {
		for j := 1; j < N; j++ {
			takeTimeConfig = append(takeTimeConfig, B)
			takeTime = append(takeTime, J)
		}
	}

	fmt.Println(takeTime)

}
