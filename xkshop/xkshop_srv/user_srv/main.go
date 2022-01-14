package main

import (
	"flag"
	"fmt"
	handler2 "xkshop/v1/xkshop_srv/user_srv/handler"
	proto2 "xkshop/v1/xkshop_srv/user_srv/proto"

	"google.golang.org/grpc"
	"net"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")
	flag.Parse()
	fmt.Println("默认ip:", *IP)
	fmt.Println("默认port:", *Port)

	server := grpc.NewServer()
	proto2.RegisterUserServer(server, &handler2.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(listen)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

}
