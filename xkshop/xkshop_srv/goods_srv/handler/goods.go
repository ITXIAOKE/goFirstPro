package handler

import (
	"gorm.io/gorm"
	"xkshop/v1/xkshop_srv/goods_srv/proto"
)

//分页功能

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

type GoodsServer struct{
	proto.UnimplementedGoodsServer
}

//商品接口

//func (s *GoodsServer) GoodsList(context.Context, *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
//
//
//}

/**
//现在用户提交订单有多个商品，你得批量查询商品的信息吧
BatchGetGoods(context.Context, *BatchGoodsIdInfo) (*GoodsListResponse, error)
CreateGoods(context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error)
DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error)
UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error)
GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error)
*/
