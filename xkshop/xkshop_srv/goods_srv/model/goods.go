package model

//定义层级极连
//实际开发过程中，尽量设置为不为null

type Category struct { //类别
	BaseModel
	Name             string    `gorm:"type:varchar(20);not null"`
	ParentCategoryID int32     //尽量使用int32  外键字段id
	ParentCategory   *Category //级连  n层面包屑
	Level            int32     `gorm:"type:int;not null;default:1"` //几级类目，默认是一级
	IsTab            bool      `gorm:"default:false;not null"`      //是否显示在首页tab
}

type Brands struct { //品牌
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct { //商品品牌分类
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

func (receiver GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct { //轮播图
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel

	CategoryID int32    `gorm:"type:int;not null"`
	Category   Category //商品类目
	BrandsID   int32    `gorm:"type:int;not null"`
	Brands     Brands   //品牌

	OnSale   bool `gorm:"default:false;not null"` //是否上架
	ShipFree bool `gorm:"default:false;not null"` //是否承担运输费
	IsNew    bool `gorm:"default:false;not null"` //是否新品
	IsHot    bool `gorm:"default:false;not null"` //是否热销

	Name     string `gorm:"type:varchar(50);not null"`   //商品名
	GoodsSn  string `gorm:"type:varchar(50);not null"`   //商品唯一货号
	ClickNum string `gorm:"type:int;default:0;not null"` //点击数
	SoldNum  string `gorm:"type:int;default:0;not null"` //商品销售量
	FavNum   string `gorm:"type:int;default:0;not null"` //商品收藏数

	MarketPrice float32 `g orm:"not null"` //市场价格
	ShopPrice   float32 `gorm:"not null"`  //本店价格

	GoodsBrief string `gorm:"type:varchar(100);not null"` //商品简短描述

	Images          GormList `gorm:"type:varchar(1000);not null"` //商品轮播图
	DescImages      GormList `gorm:"type:varchar(1000);not null"` //详情页图片
	GoodsFrontImage string   `gorm:"type:varchar(200);not null"`  //封面图片

}
