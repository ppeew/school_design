package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CollectsGetPageReq struct {
	dto.Pagination `search:"-"`
	ID             string `form:"id"`
	CollectsOrder
}

type CollectsOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:collects"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:collects"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:collects"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:collects"`
	BlogId    string `form:"blogIdOrder"  search:"type:order;column:blog_id;table:collects"`
	UserId    string `form:"userIdOrder"  search:"type:order;column:user_id;table:collects"`
}

func (m *CollectsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CollectsInsertReq struct {
	Id     int    `json:"-" comment:"主键编码"` // 主键编码
	BlogId string `json:"blogId" comment:"收藏blog ID"`
	UserId string `json:"userId" comment:"收藏者用户ID"`
	common.ControlBy
}

func (s *CollectsInsertReq) Generate(model *models.Collects) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.BlogId = s.BlogId
	model.UserId = s.UserId
}

func (s *CollectsInsertReq) GetId() interface{} {
	return s.Id
}

type CollectsUpdateReq struct {
	Id     int    `uri:"id" comment:"主键编码"` // 主键编码
	BlogId string `json:"blogId" comment:"收藏blog ID"`
	UserId string `json:"userId" comment:"收藏者用户ID"`
	common.ControlBy
}

func (s *CollectsUpdateReq) Generate(model *models.Collects) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.BlogId = s.BlogId
	model.UserId = s.UserId
}

func (s *CollectsUpdateReq) GetId() interface{} {
	return s.Id
}

// CollectsGetReq 功能获取请求参数
type CollectsGetReq struct {
	Id int `uri:"id"`
}

func (s *CollectsGetReq) GetId() interface{} {
	return s.Id
}

// CollectsDeleteReq 功能删除请求参数
type CollectsDeleteReq struct {
	Ids    []int  `json:"ids"`
	UserId string `json:"userId"`
	BlogId string `json:"blogId"`
}

func (s *CollectsDeleteReq) GetId() interface{} {
	return s.Ids
}
