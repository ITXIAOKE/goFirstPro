package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"xkshop/v1/xkshop_api/user_web/global"
	"xkshop/v1/xkshop_api/user_web/proto"
)

func InitSrvConn() {
	//从注册中心consul中获取用户服务的信息
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	userSrvHost := ""
	userSrvPort := 0

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvConfig.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn]连接【用户服务失败】")
		return
	}

	//拨号连接用户grpc服务  有一个跨域的问题options
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("【GetUserList】连接【用户服务失败】",
			"msg", err.Error(),
		)
	}

	//1，后续的问题：1，用户服务下线了2，用户服务该端口 3，用户服务该ip  在后面的负载均衡解决这个问题
	//2，好处：已经事先建立好了连接，这样后续就不用进行再次tcp的三次握手
	//3，一个连接多个Goroutine共用，性能---使用连接池
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
