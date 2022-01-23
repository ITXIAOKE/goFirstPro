package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"log"
	"xkshop/v1/xkshop_api/user_web/proto"
)

func main() {
	conn, err := grpc.Dial(
		"consul://192.168.1.101:8500/user_srv?wait=14s&tag=srv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userSrvClient := proto.NewUserClient(conn)
	for i := 0; i < 4; i++ { //测试负载均衡
		rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfoRequest{
			Pn:    1,
			PSize: 2,
		})
		if err != nil {
			panic(err)
		}

		for index, data := range rsp.Data {
			fmt.Println(index, data)
		}
	}

}
