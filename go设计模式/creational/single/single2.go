package single

import "sync"

var (
	p     *Pet
	once2 sync.Once
)

func init() {
	once2.Do(func() {
		p = &Pet{}
	})
}

func GetInstance() *Pet {
	return p
}

type Pet struct {
	name string
	age  int
	mux  sync.Mutex
}

func (p *Pet) SetName(s string) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.name = s
}

func (p *Pet) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *Pet) IncrementAge() {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.age++
}

func (p *Pet) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}
