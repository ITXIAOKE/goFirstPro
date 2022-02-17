package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	//convey.
	//	Convey("Panic recovery assertions should be accessible", t, func() {
	//	})
	//new()
	file, _ := os.Open("./aa.txt")
	fmt.Fprintf(file, "aaa")

	//myMap:=make(map[string]interface{})
	//types.Map{}

	//sync.Cond{}
	//value := atomic.Value{}
	//value.Store("aa")
	runtime.GC()

}
