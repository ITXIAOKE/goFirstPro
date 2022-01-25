package tests

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
	proto2 "xkshop/v1/xkshop_srv/goods_srv/proto"
)

var categoryClient proto2.GoodsClient
var conn3 *grpc.ClientConn
var err3 error

//测试GetAllCategorysList接口

func TestGetAllCategorysList(t *testing.T) {
	conn3, err3 = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn3.Close()
	if err3 != nil {
		panic(err3)
	}
	categoryClient = proto2.NewGoodsClient(conn3)

	rsp, err := categoryClient.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.JsonData)
	//for _, category := range rsp.Data {
	//	fmt.Println(category.Name)
	//}
}
