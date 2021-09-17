package framework

import (
	"fmt"
	"net/http"
)

type Core struct{
	name int
}


func NewCore() *Core{
	return &Core{}
}

func (c *Core) ServeHTTP(response http.ResponseWriter,request *http.Request){
	fmt.Println("world")
}