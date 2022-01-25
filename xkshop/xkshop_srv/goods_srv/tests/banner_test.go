package tests

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
	proto2 "xkshop/v1/xkshop_srv/goods_srv/proto"
)

var bannerClient proto2.GoodsClient
var conn2 *grpc.ClientConn
var err2 error

//测试BannerList接口

func TestBannerList(t *testing.T) {
	conn2, err2 = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn2.Close()
	if err2 != nil {
		panic(err2)
	}
	bannerClient = proto2.NewGoodsClient(conn2)

	rsp, err := bannerClient.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, banner := range rsp.Data {
		fmt.Println(banner.Url)
		fmt.Println(banner.Image)
	}
}
