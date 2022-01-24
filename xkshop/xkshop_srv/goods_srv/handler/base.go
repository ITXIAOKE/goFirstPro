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


type GoodsServer struct {
	proto.UnimplementedGoodsServer //为了快速测试做的
}