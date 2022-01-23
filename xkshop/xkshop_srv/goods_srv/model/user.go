package model

import (
	"gorm.io/gorm"
	"time"
)

//公共的model,每个struct都继承这个basemodel

type BaseModel struct {
	ID        int32      `gorm:"primary"`
	CreateAt  *time.Time `gorm:"column:add_time"` //所有的time都要使用指针类型的，否则不能插入数据库中，因为这个值可以为null
	UpdateAt  *time.Time `gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女 male表示男' "`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示普通用户 2表示管理员' "`
}
