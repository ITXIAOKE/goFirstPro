package proxy

//代理设计模式
import "fmt"

type Seller interface {
	sell(name string)
}

// 火车站
type Station struct {
	stock int //库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("票已售空")
	}

}

// 火车代理点
type StationProxy struct {
	station *Station // 持有一个火车站对象
}

func NewStationProxy() *StationProxy {
	return &StationProxy{station: &Station{stock: 100}}
}

func (proxy *StationProxy) sell(name string) {
	proxy.station.sell(name)
}

