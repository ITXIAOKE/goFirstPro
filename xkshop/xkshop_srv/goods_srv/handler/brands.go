package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"xkshop/v1/xkshop_srv/goods_srv/global"
	"xkshop/v1/xkshop_srv/goods_srv/model"
	"xkshop/v1/xkshop_srv/goods_srv/proto"
)

/**
//品牌和轮播图
	BrandList(context.Context, *BrandFilterRequest) (*BrandListResponse, error)
	CreateBrand(context.Context, *BrandRequest) (*BrandInfoResponse, error)
	DeleteBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
	UpdateBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
*/

//品牌列表

func (g *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {

	brandListResponse := proto.BrandListResponse{}
	var brands []model.Brands
	//brands:=make([]model.Brands,0)//以上两种方式都可以
	//result := global.DB.Find(&brands)

	//分页
	result := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var total int64
	global.DB.Model(&model.Brands{}).Count(&total) //查询出品牌表中所有的数据
	brandListResponse.Total = int32(total)

	var brandResponses []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	brandListResponse.Data = brandResponses
	return &brandListResponse, nil
}

//新建品牌

func (g *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	if result := global.DB.First(&model.Brands{}); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}
	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	global.DB.Save(brand) //更新并保存
	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}

//删除品牌

func (g *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Brands{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}

//更新品牌

func (g *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	brands := model.Brands{}
	if result := global.DB.First(&brands, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	if req.Name != "" {
		brands.Name = req.Name
	}
	if req.Logo != "" {
		brands.Logo = req.Logo
	}
	global.DB.Save(&brands)
	return &emptypb.Empty{}, nil
}
