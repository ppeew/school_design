package models

import (

	"go-admin/common/models"

)

type Comments struct {
    models.Model
    
    BlogId string `json:"blogId" gorm:"type:tinyint(4);comment:对应的博客"` 
    UserId string `json:"userId" gorm:"type:tinyint(4);comment:评论者"` 
    TargetId string `json:"targetId" gorm:"type:tinyint(4);comment:0：对作者内容评论 other：对XXX评论ID做出回复"` 
    Content string `json:"content" gorm:"type:longtext;comment:Content"` 
    models.ModelTime
    models.ControlBy
}

func (Comments) TableName() string {
    return "comments"
}

func (e *Comments) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Comments) GetId() interface{} {
	return e.Id
}