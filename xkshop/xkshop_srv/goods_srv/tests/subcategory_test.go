package tests

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	proto2 "xkshop/v1/xkshop_srv/goods_srv/proto"
)

var subcategoryClient proto2.GoodsClient

//测试GetSubCategory接口

func TestGetSubCategory(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	subcategoryClient = proto2.NewGoodsClient(conn)

	rsp, err := subcategoryClient.GetSubCategory(context.Background(), &proto2.CategoryListRequest{
		Id: 130358,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SubCategorys)
	for _, subcategory := range rsp.SubCategorys {
		fmt.Println(subcategory.Name)
	}
}
