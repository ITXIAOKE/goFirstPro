package zhuangshiqi

import (
	"log"
	"os"
	"testing"
)

//装饰器设计模式
func TestDecorator(t *testing.T) {

	f := WrapLogger(Pi, log.New(os.Stdout, "test", 1))
	f(1000000)

}
