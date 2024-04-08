package models

import "go-admin/common/models"

type Blogs struct {
	models.Model

	Username string `json:"username" gorm:"type:varchar(255);comment:用户名"`
	Title    string `json:"title" gorm:"type:varchar(255);comment:主题"`
	Msg      string `json:"msg" gorm:"type:varchar(255);comment:信息"`
	Star     string `json:"star" gorm:"type:varchar(255);comment:点赞数量"`
	Collect  string `json:"collect" gorm:"type:varchar(255);comment:收藏数量"`
	models.ModelTime
	models.ControlBy
}

func (Blogs) TableName() string {
	return "blogs"
}

func (e *Blogs) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Blogs) GetId() interface{} {
	return e.Id
}

type Comments struct {
	models.Model

	BlogId   string `json:"blogId" gorm:"type:int;comment:对应的博客"`
	UserId   string `json:"userId" gorm:"type:int;comment:评论者"`
	TargetId string `json:"targetId" gorm:"type:int;comment:0：对作者内容评论 other：对XXX评论ID做出回复"`
	models.ModelTime
	models.ControlBy
	Content string
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