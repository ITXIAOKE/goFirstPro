package tests

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	proto2 "xkshop/v1/xkshop_srv/goods_srv/proto"
)

var brandClient proto2.GoodsClient
var conn *grpc.ClientConn
var err error

//测试BrandList接口

func TestBrandList(t *testing.T) {
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	brandClient = proto2.NewGoodsClient(conn)

	rsp, err := brandClient.BrandList(context.Background(), &proto2.BrandFilterRequest{
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}
}
