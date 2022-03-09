package guanchazhe

import (
	"fmt"
	"sync"
	"time"
)

//斐波那契数列go实现

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}

type Event struct {
	Data int
}

type Observer interface {
	NotifyCallback(event Event)
}

type Subject interface {
	AddListener(obs Observer)
	RemoveListener(obs Observer)
	Notify(event Event)
}

type eventObserver struct {
	ID   int
	Time time.Time
}

type eventSubject struct {
	Observers sync.Map
}

func (e eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Recieved :%d  after : %v\n", event.Data, time.Since(e.Time))
}

func (e *eventSubject) AddListener(obs Observer) {
	e.Observers.Store(obs, struct{}{})
}

func (e *eventSubject) RemoveListener(obs Observer) {
	e.Observers.Delete(obs)
}

func (e *eventSubject) Notify(event Event) {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}
