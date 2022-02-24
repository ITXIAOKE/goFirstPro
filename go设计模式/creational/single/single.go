package single

//import "sync"
//
//type singleton behavioral {
//}
//
//var ins *singleton
//var mu sync.Mutex
//
////懒汉式
//func GetIns() *singleton {
//	if ins == nil {
//		mu.Lock()
//		if ins == nil {
//			ins = &singleton{}
//		}
//		mu.Unlock()
//	}
//	return ins
//}

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
