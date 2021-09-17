package main

import "goweb.com/ch03/framework"

func registerRouter(core *framework.Core) {
	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	core.Get("foo", FooControllerHandler)
}
