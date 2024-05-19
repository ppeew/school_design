package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CommentsGetPageReq2 struct {
	BlogID         int `form:"blogId"`
	dto.Pagination `search:"-"`
	CommentsOrder2
}

type CommentsOrder2 struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:comments"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:comments"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:comments"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:comments"`
	BlogId    string `form:"blogIdOrder"  search:"type:order;column:blog_id;table:comments"`
	UserId    string `form:"userIdOrder"  search:"type:order;column:user_id;table:comments"`
	TargetId  string `form:"targetIdOrder"  search:"type:order;column:target_id;table:comments"`
}

func (m *CommentsGetPageReq2) GetNeedSearch() interface{} {
	return *m
}

type CommentsInsertReq2 struct {
	Id       int    `json:"-" comment:"主键编码"` // 主键编码
	BlogId   string `json:"blogId" comment:"对应的博客"`
	UserId   string `json:"userId" comment:"评论者"`
	TargetId string `json:"targetId" comment:"0：对作者内容评论 other：对XXX评论ID做出回复"`
	Content  string `json:"content"`
	common.ControlBy
}

func (s *CommentsInsertReq2) Generate(model *models.Comments2) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.BlogId = s.BlogId
	model.UserId = s.UserId
	model.TargetId = s.TargetId
	model.Content = s.Content
}

func (s *CommentsInsertReq2) GetId() interface{} {
	return s.Id
}

type CommentsUpdateReq2 struct {
	Id       int    `uri:"id" comment:"主键编码"` // 主键编码
	BlogId   string `json:"blogId" comment:"对应的博客"`
	UserId   string `json:"userId" comment:"评论者"`
	TargetId string `json:"targetId" comment:"0：对作者内容评论 other：对XXX评论ID做出回复"`
	common.ControlBy
}

func (s *CommentsUpdateReq2) Generate(model *models.Comments2) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.BlogId = s.BlogId
	model.UserId = s.UserId
	model.TargetId = s.TargetId
}

func (s *CommentsUpdateReq2) GetId() interface{} {
	return s.Id
}

// CommentsGetReq2 功能获取请求参数
type CommentsGetReq2 struct {
	Id int `uri:"id"`
}

func (s *CommentsGetReq2) GetId() interface{} {
	return s.Id
}

// CommentsDeleteReq2 功能删除请求参数
type CommentsDeleteReq2 struct {
	Ids []int `json:"ids"`
}

func (s *CommentsDeleteReq2) GetId() interface{} {
	return s.Ids
}
