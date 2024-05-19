package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type Blogs2GetPageReq struct {
	dto.Pagination `search:"-"`
	BlogsOrder
}

type Blogs2Order struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:blogs"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:blogs"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:blogs"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:blogs"`
	Username  string `form:"usernameOrder"  search:"type:order;column:username;table:blogs"`
	Msg       string `form:"msgOrder"  search:"type:order;column:msg;table:blogs"`
	Star      string `form:"starOrder"  search:"type:order;column:star;table:blogs"`
	Collect   string `form:"collectOrder"  search:"type:order;column:collect;table:blogs"`
}

func (m *Blogs2GetPageReq) GetNeedSearch() interface{} {
	return *m
}

type Blogs2InsertReq struct {
	Id       int    `json:"-" comment:"主键编码"` // 主键编码
	Username string `json:"username" comment:"用户名"`
	Msg      string `json:"msg" comment:"信息"`
	Star     string `json:"star" comment:"点赞数量"`
	Collect  string `json:"collect" comment:"收藏数量"`
	Title    string `json:"title"`
	common.ControlBy
}

func (s *Blogs2InsertReq) Generate(model *models.Blogs) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Msg = s.Msg
	model.Star = s.Star
	model.Collect = s.Collect
	model.Title = s.Title
}

func (s *Blogs2InsertReq) GetId() interface{} {
	return s.Id
}

type Blogs2UpdateReq struct {
	Type     string `form:"type"`
	Id       int    `uri:"id" comment:"主键编码"` // 主键编码
	Username string `json:"username" comment:"用户名"`
	Msg      string `json:"msg" comment:"信息"`
	Star     string `json:"star" comment:"点赞数量"`
	Collect  string `json:"collect" comment:"收藏数量"`
	common.ControlBy
}

func (s *Blogs2UpdateReq) Generate(model *models.Blogs) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Msg = s.Msg
	model.Star = s.Star
	model.Collect = s.Collect
}

func (s *Blogs2UpdateReq) GetId() interface{} {
	return s.Id
}

// BlogsGetReq 功能获取请求参数
type Blogs2GetReq struct {
	Id int `uri:"id"`
}

func (s *Blogs2GetReq) GetId() interface{} {
	return s.Id
}

// BlogsDeleteReq 功能删除请求参数
type Blogs2DeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *Blogs2DeleteReq) GetId() interface{} {
	return s.Ids
}
