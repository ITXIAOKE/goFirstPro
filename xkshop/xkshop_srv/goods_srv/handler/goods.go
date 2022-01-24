package handler

import (
	"context"
	"xkshop/v1/xkshop_srv/goods_srv/proto"
)

//商品接口

func (g *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {

	return nil, nil

}

/**
//现在用户提交订单有多个商品，你得批量查询商品的信息吧
BatchGetGoods(context.Context, *BatchGoodsIdInfo) (*GoodsListResponse, error)
CreateGoods(context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error)
DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error)
UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error)
GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error)
*/
