package tests

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	proto2 "xkshop/v1/xkshop_srv/goods_srv/proto"
)

var brandcategoryClient proto2.GoodsClient

//测试GetSubCategory接口

func TestGetCategoryBrandList(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	brandcategoryClient = proto2.NewGoodsClient(conn)

	rsp, err := brandcategoryClient.GetCategoryBrandList(context.Background(), &proto2.CategoryInfoRequest{
		Id: int32(136842),
		Level: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, data := range rsp.Data {
		fmt.Println(data.Name)
	}
}


func TestCategoryBrandList(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	brandcategoryClient = proto2.NewGoodsClient(conn)

	rsp, err := brandcategoryClient.CategoryBrandList(context.Background(), &proto2.CategoryBrandFilterRequest{
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)

}
