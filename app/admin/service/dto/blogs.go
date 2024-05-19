package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BlogsGetPageReq struct {
	dto.Pagination     `search:"-"`
    BlogsOrder
}

type BlogsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:blogs"`
    Username string `form:"usernameOrder"  search:"type:order;column:username;table:blogs"`
    Title string `form:"titleOrder"  search:"type:order;column:title;table:blogs"`
    Msg string `form:"msgOrder"  search:"type:order;column:msg;table:blogs"`
    Star string `form:"starOrder"  search:"type:order;column:star;table:blogs"`
    Collect string `form:"collectOrder"  search:"type:order;column:collect;table:blogs"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:blogs"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:blogs"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:blogs"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:blogs"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:blogs"`
    
}

func (m *BlogsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BlogsInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    Username string `json:"username" comment:"用户名"`
    Title string `json:"title" comment:"主题"`
    Msg string `json:"msg" comment:"信息"`
    Star string `json:"star" comment:"点赞数量"`
    Collect string `json:"collect" comment:"收藏数量"`
    common.ControlBy
}

func (s *BlogsInsertReq) Generate(model *models.Blogs)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Username = s.Username
    model.Title = s.Title
    model.Msg = s.Msg
    model.Star = s.Star
    model.Collect = s.Collect
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BlogsInsertReq) GetId() interface{} {
	return s.Id
}

type BlogsUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Username string `json:"username" comment:"用户名"`
    Title string `json:"title" comment:"主题"`
    Msg string `json:"msg" comment:"信息"`
    Star string `json:"star" comment:"点赞数量"`
    Collect string `json:"collect" comment:"收藏数量"`
    common.ControlBy
}

func (s *BlogsUpdateReq) Generate(model *models.Blogs)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Username = s.Username
    model.Title = s.Title
    model.Msg = s.Msg
    model.Star = s.Star
    model.Collect = s.Collect
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BlogsUpdateReq) GetId() interface{} {
	return s.Id
}

// BlogsGetReq 功能获取请求参数
type BlogsGetReq struct {
     Id int `uri:"id"`
}
func (s *BlogsGetReq) GetId() interface{} {
	return s.Id
}

// BlogsDeleteReq 功能删除请求参数
type BlogsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BlogsDeleteReq) GetId() interface{} {
	return s.Ids
}
