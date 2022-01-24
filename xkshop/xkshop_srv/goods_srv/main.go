package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
	"xkshop/v1/xkshop_api/user_web/utils"
	"xkshop/v1/xkshop_srv/goods_srv/global"
	handler2 "xkshop/v1/xkshop_srv/goods_srv/handler"
	"xkshop/v1/xkshop_srv/goods_srv/initialize"
	proto2 "xkshop/v1/xkshop_srv/goods_srv/proto"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	//Port := flag.Int("port", 0, "端口号")//生产用这种
	Port := flag.Int("port", 50051, "端口号") //开发用这种固定的端口，方便测试

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	zap.S().Info(global.ServerConfig)

	flag.Parse()

	zap.S().Info("默认ip:", *IP)
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Info("默认port:", *Port)

	server := grpc.NewServer()
	proto2.RegisterGoodsServer(server, &handler2.GoodsServer{})
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
		GRPC:                           fmt.Sprintf("%s:%d", global.ServerConfig.Host, *Port), //这里必须是这个格式才生效
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s", //注销
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4()) //id不同就可以生成多个consul实例
	registration.ID = serviceID
	registration.Port = *Port
	registration.Tags = global.ServerConfig.Tags
	registration.Address = global.ServerConfig.Host
	registration.Check = check

	err = client.Agent().ServiceRegister(registration) //注册服务
	//client.Agent().ServiceDeregister(name)//撤销服务
	if err != nil {
		panic(err)
	}

	go func() {
		err = server.Serve(listen) //这个是阻塞的，必须放在goroutine中，否则后面的代码无法执行
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接收终止信号   优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
