package single

import (
	"sync"
)

type singleton struct {
}

var ins *singleton
var once sync.Once

//更符合go设计

func GetInsOr() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
