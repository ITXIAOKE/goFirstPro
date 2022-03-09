package zhuangshiqi

import (
	"log"
	"math"
	"time"
)

type piFunc func(int) float64

func WrapLogger(fun piFunc, logger *log.Logger) piFunc {
	return func(i int) float64 {
		fn := func(i int) (result float64) {
			defer func(t time.Time) {
				logger.Printf("took=%v, i=%v, result=%v", time.Since(t), i, result)
			}(time.Now())

			return fun(i)
		}

		return fn(i)
	}
}

func Pi(n int) float64 {
	ch := make(chan float64)
	for i := 0; i < n; i++ {
		go func(ch chan float64, i float64) {
			ch <- 4 * math.Pow(-1, i) / (2*i + 1)
		}(ch, float64(i))
	}
	result := 0.0 //这个地方一定要0.0表示是float类型，不能是0就成int类型，后面就报错
	for i := 0; i < n; i++ {
		result += <-ch
	}
	return result
}
