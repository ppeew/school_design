package apis

import (
	"fmt"
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

type Collects struct {
	api.Api
}

// GetPage 获取Collects列表
// @Summary 获取Collects列表
// @Description 获取Collects列表
// @Tags Collects
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Collects}} "{"code": 200, "data": [...]}"
// @Router /api/v1/collects [get]
// @Security Bearer
func (e Collects) GetPage(c *gin.Context) {
	req := dto.CollectsGetPageReq{}
	s := service.Collects{}
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
	list := make([]models.Collects, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Collects失败，\r\n失败信息 %s", err.Error()))
		return
	}

	type ss struct {
		BlogID      int
		Username    string
		BlogTime    string
		Image       string
		BlogTitle   string
		BlogComment int
		BlogCollect int
		CollectTime string
	}

	resp := make([]ss, 0)
	for _, collects := range list {
		u := new(models.User)
		e.Orm.Where("id=?", collects.UserId).First(u)
		blog := new(models.Blogs)
		e.Orm.Where("id=?", collects.BlogId).First(blog)
		var commentTotal int64
		e.Orm.Model(models.Comments{}).Where("blog_id=?", collects.BlogId).Count(&commentTotal)
		var collectTotal int64
		e.Orm.Model(models.Collects{}).Where("blog_id=?", collects.BlogId).Count(&collectTotal)
		if u.Image == "" {
			u.Image = "https://th.bing.com/th/id/R.ec214a16a4b823260966fdb7b09eddff?rik=ivXWTf6iBDSAsg&riu=http%3a%2f%2fpic2.nipic.com%2f20090429%2f984755_194557035_2.jpg&ehk=u7%2bUd79GB72em%2b7h%2bdvk6exyLNUKzpSdUpeA8DWRn2s%3d&risl=&pid=ImgRaw&r=0"
		}

		resp = append(resp, ss{
			BlogID:      blog.Id,
			Username:    u.Username,
			BlogTime:    blog.CreatedAt.Format(time.DateTime),
			Image:       u.Image,
			BlogTitle:   blog.Title,
			BlogComment: int(commentTotal),
			BlogCollect: int(collectTotal),
			CollectTime: collects.CreatedAt.Format(time.DateTime),
		})
	}

	e.PageOK(resp, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Collects
// @Summary 获取Collects
// @Description 获取Collects
// @Tags Collects
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Collects} "{"code": 200, "data": [...]}"
// @Router /api/v1/collects/{id} [get]
// @Security Bearer
func (e Collects) Get(c *gin.Context) {
	req := dto.CollectsGetReq{}
	s := service.Collects{}
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
	var object models.Collects

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Collects失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Collects
// @Summary 创建Collects
// @Description 创建Collects
// @Tags Collects
// @Accept application/json
// @Product application/json
// @Param data body dto.CollectsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/collects [post]
// @Security Bearer
func (e Collects) Insert(c *gin.Context) {
	req := dto.CollectsInsertReq{}
	s := service.Collects{}
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
		e.Error(500, err, fmt.Sprintf("创建Collects失败，\r\n失败信息 %s", err.Error()))
		return
	}

	me := models.Message{}
	me.Type = 2
	me.NoticeStatus = 0
	u := new(models.User)
	e.Orm.Where("id=?", req.UserId).First(u)
	bl := new(models.Blogs)
	e.Orm.Where("id=?", req.BlogId).First(bl)
	me.Content = fmt.Sprintf("%s收藏了你的博客：%s", u.Username, bl.Msg)
	u2 := new(models.User)
	e.Orm.Where("username=?", bl.Username).First(u2)
	me.UserID = u2.Id
	e.Orm.Create(&me)

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Collects
// @Summary 修改Collects
// @Description 修改Collects
// @Tags Collects
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.CollectsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/collects/{id} [put]
// @Security Bearer
func (e Collects) Update(c *gin.Context) {
	req := dto.CollectsUpdateReq{}
	s := service.Collects{}
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
		e.Error(500, err, fmt.Sprintf("修改Collects失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Collects
// @Summary 删除Collects
// @Description 删除Collects
// @Tags Collects
// @Param data body dto.CollectsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/collects [delete]
// @Security Bearer
func (e Collects) Delete(c *gin.Context) {
	s := service.Collects{}
	req := dto.CollectsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Collects失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

func (e Collects) RemoveCollect(c *gin.Context) {
	s := service.Collects{}
	req := dto.CollectsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Collects失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK("", "删除成功")
}
