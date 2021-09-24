package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string   //使用切片来模拟队列，存数据
	cond  *sync.Cond //让一组goroutine在满足特定条件时被唤醒
}

func (q *Queue) Enqueue(item string) { //生产者
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Printf("putting %s to queue, notify all\n", item)
	q.cond.Broadcast() //通知所有的goroutine
	//q.cond.Signal() //随机的通知一个goroutine
}

func (q *Queue) Dequeue() string { //消费者
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result //把队列的头元素给返回
}

func main() {
	/**
	使用cond实现生产者和消费者
	*/
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	go func() {
		for {
			q.Enqueue("aadfadfa")
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		fmt.Println("消费的数据是：" + q.Dequeue())
		time.Sleep(time.Second)
	}
}
