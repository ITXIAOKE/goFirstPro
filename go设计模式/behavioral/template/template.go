package template

import "fmt"

type WorkInterface interface {
	GetUp()
	Work()
	Sleep()
}

type Worker struct {
	WorkInterface
}

func NewWork(w WorkInterface) *Worker {
	return &Worker{
		w,
	}
}

func (w *Worker) Daily() {
	w.GetUp()
	w.Work()
	w.Sleep()
}

type Corder struct {
}

func (c *Corder) GetUp() {
	fmt.Println("coder getup..")
}

func (c *Corder) Work() {
	fmt.Println("coder work...")
}

func (c *Corder) Sleep() {
	fmt.Println("coder sleep..")
}
