package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	"google.golang.org/grpc"

	"xkshop/v1/grpc_project/grpc_error_test/proto"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {
	return nil, status.Errorf(codes.NotFound, "记录未找到：%s", request.Name)
	//return &proto.HelloReply{
	//	Message: "hello "+request.Name,
	//}, nil
}

func main() {
	/**
	server端：最好是帮我们生成接口，我们只需要去每个接口中实现对象的业务逻辑就行了
	client端：需要帮我们生成对应的方法，同时将这个方法都邦定到一个结构体上，生成的时候我们可能需要传参数，ip:port
	*/
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
