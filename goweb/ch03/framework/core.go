package framework

import (
	"log"
	"net/http"
)

// 框架核心结构

type Core struct {
	router map[string]ControllerHandler
}

// 初始化框架核心结构

func NewCore() *Core {
	return &Core{router:map[string]ControllerHandler{}}
}

func (c *Core) Get(url string,handler ControllerHandler){
	c.router[url]=handler
}

func (c *Core) Post(url string,handler ControllerHandler){
	c.router[url]=handler
}

func (c *Core) FindRouterByRequest(request *http.Request) ControllerHandler{
	return c.router["foo"]
}



// 框架核心结构实现Handler接口

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serverHTTP")
	ctx:=NewContext(request,response)

	router:=c.FindRouterByRequest(request)
	if router == nil {
		return
	}
	log.Println("core.router")
	ctx.SetHandler(router)
	router(ctx)

}
