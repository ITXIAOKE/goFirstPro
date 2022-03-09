package single

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingles(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			MyIncrementAge()
		}()

		go func() {
			defer wg.Done()
			MyIncrementAge2()
		}()
	}
	wg.Wait()
	p := GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
