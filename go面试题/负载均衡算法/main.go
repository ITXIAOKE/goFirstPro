//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//type LoadBalancer behavioral {
//	client []*Client
//	size int32
//}
//
//func NewLoadBalancer(size int32) *LoadBalancer{
//	loadBalance:=&LoadBalancer{client: make([]*Client,size),size: size}
//	loadBalance.client=append(loadBalance.client,&Client{})
//	return loadBalance
//}
//
//func (m *LoadBalancer)getClient() *Client{
//	rand.Seed(time.Now().Unix())
//	x:=rand.Int31n(100)
//	return m.client[x%m.size]
//}
//
//type Client behavioral {
//	Name string
//}
//
//func (c *Client) Do()  {
//	fmt.Println("do")
//}
//
//func main() {
//	lb:=NewLoadBalancer(40)
//	lb.getClient().Do()
//
//}

package main

import (
	"fmt"
)

type RoundRobin struct {
	servers []string
	current int
}

/** 获取下一个服务器 */
func (R *RoundRobin) next() string {
	R.current++
	R.current = R.current % len(R.servers) // 访问到最后一个服务器之后，重置会第一台。 5%5=0。
	return R.servers[R.current]
}

func main() {
	r := &RoundRobin{
		servers: []string{"192.168.10", "192.168.11", "192.168.12"},
		current: -1,
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("| %d | %s |\n", i+1, r.next())
	}
}
