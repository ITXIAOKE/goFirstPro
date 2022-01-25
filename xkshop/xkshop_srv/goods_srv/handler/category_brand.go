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
//品牌分类   选择完分类后，则对应的品牌就应展示出来

	CategoryBrandList(context.Context, *CategoryBrandFilterRequest) (*CategoryBrandListResponse, error)
	//通过category获取brands
	GetCategoryBrandList(context.Context, *CategoryInfoRequest) (*BrandListResponse, error)
	CreateCategoryBrand(context.Context, *CategoryBrandRequest) (*CategoryBrandResponse, error)
	DeleteCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error)
	UpdateCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error)
*/

//获取全部品牌分类

func (g *GoodsServer) CategoryBrandList(ctx context.Context, req *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	var categoryBrands []model.GoodsCategoryBrand
	categoryBrandsListResponse := proto.CategoryBrandListResponse{}

	var total int64
	global.DB.Model(&model.GoodsCategoryBrand{}).Count(&total)
	categoryBrandsListResponse.Total = int32(total)

	//global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)//数据都为空，因为有外键
	global.DB.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)

	var categoryResponses []*proto.CategoryBrandResponse
	for _, categoryBrand := range categoryBrands {
		categoryResponses = append(categoryResponses, &proto.CategoryBrandResponse{
			Category: &proto.CategoryInfoResponse{
				Id:             categoryBrand.Category.ID,
				Name:           categoryBrand.Category.Name,
				Level:          categoryBrand.Category.Level,
				IsTab:          categoryBrand.Category.IsTab,
				ParentCategory: categoryBrand.Category.ParentCategoryID,
			},
			Brand: &proto.BrandInfoResponse{
				Id:   categoryBrand.Brands.ID,
				Name: categoryBrand.Brands.Name,
				Logo: categoryBrand.Brands.Logo,
			},
		})
	}

	categoryBrandsListResponse.Data = categoryResponses
	return &categoryBrandsListResponse, nil

}

//通过category获取brands

func (g *GoodsServer) GetCategoryBrandList(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {

	brandListResponse := proto.BrandListResponse{}

	var category model.Category
	if result := global.DB.Find(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var categoryBrands []model.GoodsCategoryBrand
	if result := global.DB.Preload("Brands").Where(&model.GoodsCategoryBrand{CategoryID: req.Id}).Find(&categoryBrands); result.RowsAffected > 0 {
		brandListResponse.Total = int32(result.RowsAffected)
	}

	var brandInfoResponses []*proto.BrandInfoResponse
	for _, categoryBrand := range categoryBrands {
		brandInfoResponses = append(brandInfoResponses, &proto.BrandInfoResponse{
			Id:   categoryBrand.Brands.ID,
			Name: categoryBrand.Brands.Name,
			Logo: categoryBrand.Brands.Logo,
		})
	}

	brandListResponse.Data = brandInfoResponses
	return &brandListResponse, nil

}

func (g *GoodsServer) CreateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand := model.GoodsCategoryBrand{
		CategoryID: req.CategoryId,
		BrandsID:   req.BrandId,
	}

	global.DB.Save(&categoryBrand)
	return &proto.CategoryBrandResponse{Id: categoryBrand.ID}, nil

}

func (g *GoodsServer) DeleteCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.GoodsCategoryBrand{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}
	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {

	var categoryBrand model.GoodsCategoryBrand
	if result := global.DB.First(&categoryBrand, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌分类不存在")
	}

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand.CategoryID = req.CategoryId
	categoryBrand.BrandsID = req.BrandId

	global.DB.Save(&categoryBrand)
	return &emptypb.Empty{}, nil
}
