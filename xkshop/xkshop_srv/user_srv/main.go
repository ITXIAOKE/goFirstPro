package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"xkshop/v1/xkshop_srv/user_srv/global"
	handler2 "xkshop/v1/xkshop_srv/user_srv/handler"
	"xkshop/v1/xkshop_srv/user_srv/initialize"
	proto2 "xkshop/v1/xkshop_srv/user_srv/proto"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	zap.S().Info(global.ServerConfig)

	flag.Parse()

	zap.S().Info("默认ip:", *IP)
	zap.S().Info("默认port:", *Port)

	server := grpc.NewServer()
	proto2.RegisterUserServer(server, &handler2.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	//注册grpc服务健康检查到consul中
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("192.168.1.102:50051"), //这里必须是这个格式才生效
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = *Port
	registration.Tags = []string{"xkshop", "xiaoke", "user", "srv"}
	registration.Address = "192.168.1.102"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration) //注册服务
	//client.Agent().ServiceDeregister(name)//撤销服务
	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

}
