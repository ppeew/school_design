package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CommentsGetPageReq struct {
	dto.Pagination     `search:"-"`
    CommentsOrder
}

type CommentsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:comments"`
    BlogId string `form:"blogIdOrder"  search:"type:order;column:blog_id;table:comments"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:comments"`
    TargetId string `form:"targetIdOrder"  search:"type:order;column:target_id;table:comments"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:comments"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:comments"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:comments"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:comments"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:comments"`
    Content string `form:"contentOrder"  search:"type:order;column:content;table:comments"`
    
}

func (m *CommentsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CommentsInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    BlogId string `json:"blogId" comment:"对应的博客"`
    UserId string `json:"userId" comment:"评论者"`
    TargetId string `json:"targetId" comment:"0：对作者内容评论 other：对XXX评论ID做出回复"`
    Content string `json:"content" comment:""`
    common.ControlBy
}

func (s *CommentsInsertReq) Generate(model *models.Comments)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BlogId = s.BlogId
    model.UserId = s.UserId
    model.TargetId = s.TargetId
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.Content = s.Content
}

func (s *CommentsInsertReq) GetId() interface{} {
	return s.Id
}

type CommentsUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    BlogId string `json:"blogId" comment:"对应的博客"`
    UserId string `json:"userId" comment:"评论者"`
    TargetId string `json:"targetId" comment:"0：对作者内容评论 other：对XXX评论ID做出回复"`
    Content string `json:"content" comment:""`
    common.ControlBy
}

func (s *CommentsUpdateReq) Generate(model *models.Comments)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BlogId = s.BlogId
    model.UserId = s.UserId
    model.TargetId = s.TargetId
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.Content = s.Content
}

func (s *CommentsUpdateReq) GetId() interface{} {
	return s.Id
}

// CommentsGetReq 功能获取请求参数
type CommentsGetReq struct {
     Id int `uri:"id"`
}
func (s *CommentsGetReq) GetId() interface{} {
	return s.Id
}

// CommentsDeleteReq 功能删除请求参数
type CommentsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *CommentsDeleteReq) GetId() interface{} {
	return s.Ids
}
