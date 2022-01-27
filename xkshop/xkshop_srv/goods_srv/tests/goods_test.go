package tests

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	proto2 "xkshop/v1/xkshop_srv/goods_srv/proto"
)

var goodsListClient proto2.GoodsClient

//测试GetSubCategory接口

func TestGoodsList(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	goodsListClient = proto2.NewGoodsClient(conn)

	rsp, err := goodsListClient.GoodsList(context.Background(), &proto2.GoodsFilterRequest{
		TopCategory: 130361,
		PriceMin:    90,
		//KeyWords: "国产海南冷冻鲷鱼柳",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)

	for _, data := range rsp.Data {
		fmt.Println(data.Name, data.ShopPrice)
	}
}

var goodsbatchClient proto2.GoodsClient

func TestBatchGetGoods(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	goodsbatchClient = proto2.NewGoodsClient(conn)

	rsp, err := goodsbatchClient.BatchGetGoods(context.Background(), &proto2.BatchGoodsIdInfo{
		Id: []int32{421, 422, 423},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)

	for _, data := range rsp.Data {
		fmt.Println(data.Name, data.ShopPrice)
	}
}

var goodsdetialClient proto2.GoodsClient

func TestGetGoodsDetail(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	goodsdetialClient = proto2.NewGoodsClient(conn)

	data, err := goodsdetialClient.GetGoodsDetail(context.Background(), &proto2.GoodInfoRequest{
		Id: 421,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Name, data.ShopPrice)
}
