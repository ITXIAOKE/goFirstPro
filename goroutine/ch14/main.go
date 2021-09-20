package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	RandRum int
}

type Result struct {
	job *Job
	sum int
}

func main() {
	/**
		需求：
	计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6 随机生成数字进行计算
	*/
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	createPool(2, jobChan, resultChan)

	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandRum, result.sum)
		}
	}(resultChan)

	//主goroutine，循环创建job，输入到管道
	var id int
	for {
		id++
		randNum := rand.Int()
		job := &Job{
			Id:      id,
			RandRum: randNum,
		}
		jobChan <- job
	}

}

func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	//根据开协程个数，去跑几个goroutine
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				randNum := job.RandRum
				var sum int
				for randNum != 0 {
					temp := randNum % 10
					sum += temp
					randNum /= 10
				}
				re := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- re
			}
		}(jobChan, resultChan)
	}
}
