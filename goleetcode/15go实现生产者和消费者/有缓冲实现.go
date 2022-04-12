package _5go实现生产者和消费者

//在这个程序中，缓冲区可以存储10个int类型的整数，在执行生产者线程的时候，
//线程就不会阻塞，一次性将10个整数存入channel，在读取的时候，也是一次性读取。


// 带缓冲区的channel
import (
	"fmt"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Send:", i)
	}
}

func consumer(ch <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("Receive:", v)
	}
}

func main() {
	ch := make(chan int, 10)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}