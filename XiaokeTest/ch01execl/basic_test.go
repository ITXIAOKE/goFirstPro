package main

import (
	"testing"
)

//表格驱动测试
func TestTriangle(t *testing.T) { //覆盖率
	/**
	go test
	go test -coverprofile=c.out
	go tool cover
	go tool cover -html=c.out

	go test -bench . -cpuprofile cpu.out
	go tool pprof cpu.out
	*/
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); "+"got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}

}

/**
性能调优的4个步骤：
1，go test -bench . -cpuprofile cpu.out
2，go tool pprof cpu.out
3，分析慢在哪里？那个框大，那里耗时就多
4，优化代码

godoc -http :6060

在测试文件中加入Example生成实例代码：
func ExampleQueue(){}
*/
func BenchmarkTriangle(b *testing.B) { // 性能测试
	tests := []struct{ a, b, c int }{
		{300000, 400000, 500000},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
				b.Errorf("calcTriangle(%d, %d); "+"got %d; expected %d", tt.a, tt.b, actual, tt.c)
			}
		}
	}

}
