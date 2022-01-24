package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

//自定义数据类型

type GormList []string

func (receiver GormList) Value() (driver.Value, error) {
	return json.Marshal(receiver)
}

func (receiver *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &receiver)
}

//公共的model,每个struct都继承这个basemodel

type BaseModel struct {
	ID        int32      `gorm:"primarykey;type:int"`
	CreateAt  *time.Time `gorm:"column:add_time"` //所有的time都要使用指针类型的，否则不能插入数据库中，因为这个值可以为null
	UpdateAt  *time.Time `gorm:"column:update_time"`
	//DeletedAt  gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}
