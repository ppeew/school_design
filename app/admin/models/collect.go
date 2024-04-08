package models

import "go-admin/common/models"

type Collects struct {
	models.Model

	BlogId string `json:"blogId" gorm:"type:int;comment:收藏blog ID"`
	UserId string `json:"userId" gorm:"type:int;comment:收藏者用户ID"`
	models.ModelTime
	models.ControlBy
}

func (Collects) TableName() string {
	return "collects"
}

func (e *Collects) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Collects) GetId() interface{} {
	return e.Id
}
