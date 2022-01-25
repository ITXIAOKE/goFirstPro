package handler

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"xkshop/v1/xkshop_srv/goods_srv/global"
	"xkshop/v1/xkshop_srv/goods_srv/model"
	"xkshop/v1/xkshop_srv/goods_srv/proto"
)

/**
//商品分类
GetAllCategorysList(context.Context, *emptypb.Empty) (*CategoryListResponse, error)
//获取子分类
GetSubCategory(context.Context, *CategoryListRequest) (*SubCategoryListResponse, error)
CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error)
*/

//商品分类

func (g *GoodsServer) GetAllCategorysList(ctx context.Context, req *emptypb.Empty) (*proto.CategoryListResponse, error) {

	var categorys []model.Category

	//global.DB.Preload("SubCategory").Find(&categorys) //这种不能拿到三级目录数据
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	// SELECT * FROM `category` WHERE `category`.`parent_category_id` IN (130358,130361,130364,130365
	// SELECT * FROM `category` WHERE `category`.`level` = 1

	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

//获取子分类,不用分页就行，全部取出，对照京东的首页，面包协分类

func (g *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	subCategoryListResponse := proto.SubCategoryListResponse{}
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	subCategoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategorys []model.Category
	var subCategoryResponse []*proto.CategoryInfoResponse
	preloads := "SubCategory"
	if category.Level == 1 {
		preloads = "SubCategory.SubCategory"
	}
	global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Preload(preloads).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		subCategoryResponse = append(subCategoryResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}
	subCategoryListResponse.SubCategorys = subCategoryResponse
	return &subCategoryListResponse, nil
}

//新建分类

func (g *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {

	category := model.Category{}

	category.Name = req.Name
	category.Level = req.Level

	if req.Level != 1 {
		category.ParentCategoryID = req.ParentCategory
	}
	category.IsTab = req.IsTab

	global.DB.Save(&category)

	return &proto.CategoryInfoResponse{Id: int32(category.ID)}, nil

}

//删除分类

func (g *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &emptypb.Empty{}, nil
}

//更新商品分类

func (g *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {

	var category model.Category

	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}

	if req.Level != 0 {
		category.Level = req.Level
	}

	if req.IsTab {
		category.IsTab = req.IsTab
	}

	global.DB.Save(&category)

	return &emptypb.Empty{}, nil
}
