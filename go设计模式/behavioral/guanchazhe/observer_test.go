package guanchazhe

import (
	"sync"
	"testing"
	"time"
)

//观察者设计模式

func TestObserver(t *testing.T) {
	//for i2 := range Fib(10){
	//	fmt.Println(i2)
	//}

	n := eventSubject{Observers: sync.Map{}} //主题

	obs1 := eventObserver{ID: 1, Time: time.Now()} //观察者
	obs2 := eventObserver{ID: 2, Time: time.Now()}

	n.AddListener(obs1)
	n.AddListener(obs2)

	for x := range Fib(10) {
		n.Notify(Event{Data: x})
	}

}
