package proxy

import "testing"

func TestProxy(t *testing.T) {
	proxy := NewStationProxy()
	proxy.sell("昌平火车站")

}
