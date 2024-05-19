package models

import (
	"go-admin/common/models"
)

type User2 struct {
	models.Model

	Username string `json:"username" gorm:"type:varchar(255);comment:用户名"`
	Password string `json:"password" gorm:"type:varchar(255);comment:密码"`
	Age      string `json:"age" gorm:"type:bigint;comment:年龄"`
	Sex      string `json:"sex" gorm:"type:bigint;comment:性别,1:男 2:女"`
	models.ModelTime
	models.ControlBy
	Image string `json:"image" gorm:"type:varchar(255);comment:用户图像url"`
	Role  int    `json:"role" gorm:"type:int;comment:职责，1：管理员 0：普通用户"`
}

func (User2) TableName() string {
	return "user"
}

func (e *User2) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *User2) GetId() interface{} {
	return e.Id
}
