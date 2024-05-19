package apis

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Comments2 struct {
	api.Api
}

// GetPage 获取Comments列表
// @Summary 获取Comments列表
// @Description 获取Comments列表
// @Tags Comments2
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Comments2}} "{"code": 200, "data": [...]}"
// @Router /api/v1/comments [get]
// @Security Bearer
func (e Comments2) GetPage(c *gin.Context) {
	req := dto.CommentsGetPageReq2{}
	s := service.Comments2{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Comments2, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Comments失败，\r\n失败信息 %s", err.Error()))
		return
	}

	type bbb struct {
		SecCommentUserName      string
		SecCommentUserAvatarUrl string
		SecContent              string
		CreatedTime             string
	}

	type aaa struct {
		ID          int
		UserImage   string
		UserID      int
		CommentTime string
		Content     string
		UserName    string

		// 评论回复
		CommentResp []bbb
	}

	ret := make([]aaa, 0)
	for _, comments := range list {
		user := new(models.User2)
		e.Orm.Model(&models.User2{}).Where("id=?", comments.UserId).First(user)

		commentResp := make([]bbb, 0)
		//查询该评论对应的回复
		c2 := make([]models.Comments2, 0)
		e.Orm.Model(&models.Comments2{}).Where("blog_id=?", req.BlogID).Where("target_id=?", comments.Id).Find(&c2)
		for _, m := range c2 {
			user2 := new(models.User2)
			e.Orm.Model(&models.User2{}).Where("id=?", m.UserId).First(user2)
			commentResp = append(commentResp, bbb{
				SecCommentUserName:      user2.Username,
				SecCommentUserAvatarUrl: user2.Image,
				SecContent:              m.Content,
				CreatedTime:             m.CreatedAt.Format(time.DateTime),
			})
		}

		ret = append(ret, aaa{
			ID:          comments.Id,
			UserImage:   user.Image,
			UserID:      user.Id,
			UserName:    user.Username,
			CommentTime: comments.CreatedAt.Format(time.DateTime),
			Content:     comments.Content,
			CommentResp: commentResp,
		})
	}

	e.PageOK(ret, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Comments
// @Summary 获取Comments
// @Description 获取Comments
// @Tags Comments2
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Comments2} "{"code": 200, "data": [...]}"
// @Router /api/v1/comments/{id} [get]
// @Security Bearer
func (e Comments2) Get(c *gin.Context) {
	req := dto.CommentsGetReq2{}
	s := service.Comments2{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Comments2

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Comments失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Comments
// @Summary 创建Comments
// @Description 创建Comments
// @Tags Comments2
// @Accept application/json
// @Product application/json
// @Param data body dto.CommentsInsertReq2 true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/comments [post]
// @Security Bearer
func (e Comments2) Insert(c *gin.Context) {
	req := dto.CommentsInsertReq2{}
	s := service.Comments2{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建Comments失败，\r\n失败信息 %s", err.Error()))
		return
	}

	me := models.Message{}
	if req.TargetId == "0" {
		blog := models.Blogs{}
		e.Orm.Where("id=?", req.BlogId).First(&blog)
		u := new(models.User2)
		e.Orm.Where("username=?", blog.Username).First(u)
		me.UserID = u.Id
	} else {
		me.UserID = gconv.Int(req.TargetId)
	}
	me.Type = 3
	me.NoticeStatus = 0

	u := new(models.User2)
	e.Orm.Where("id=?", req.UserId).First(u)
	me.Content = fmt.Sprintf("%s对你说：%s", u.Username, req.Content)

	e.Orm.Create(&me)

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Comments
// @Summary 修改Comments
// @Description 修改Comments
// @Tags Comments2
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.CommentsUpdateReq2 true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/comments/{id} [put]
// @Security Bearer
func (e Comments2) Update(c *gin.Context) {
	req := dto.CommentsUpdateReq2{}
	s := service.Comments2{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改Comments失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Comments
// @Summary 删除Comments
// @Description 删除Comments
// @Tags Comments2
// @Param data body dto.CommentsDeleteReq2 true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/comments [delete]
// @Security Bearer
func (e Comments2) Delete(c *gin.Context) {
	s := service.Comments2{}
	req := dto.CommentsDeleteReq2{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除Comments失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
