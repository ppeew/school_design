package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Comments struct {
	api.Api
}

// GetPage 获取Comments列表
// @Summary 获取Comments列表
// @Description 获取Comments列表
// @Tags Comments
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Comments}} "{"code": 200, "data": [...]}"
// @Router /api/v1/comments [get]
// @Security Bearer
func (e Comments) GetPage(c *gin.Context) {
    req := dto.CommentsGetPageReq{}
    s := service.Comments{}
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
	list := make([]models.Comments, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Comments失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Comments
// @Summary 获取Comments
// @Description 获取Comments
// @Tags Comments
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Comments} "{"code": 200, "data": [...]}"
// @Router /api/v1/comments/{id} [get]
// @Security Bearer
func (e Comments) Get(c *gin.Context) {
	req := dto.CommentsGetReq{}
	s := service.Comments{}
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
	var object models.Comments

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Comments失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Comments
// @Summary 创建Comments
// @Description 创建Comments
// @Tags Comments
// @Accept application/json
// @Product application/json
// @Param data body dto.CommentsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/comments [post]
// @Security Bearer
func (e Comments) Insert(c *gin.Context) {
    req := dto.CommentsInsertReq{}
    s := service.Comments{}
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

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Comments
// @Summary 修改Comments
// @Description 修改Comments
// @Tags Comments
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.CommentsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/comments/{id} [put]
// @Security Bearer
func (e Comments) Update(c *gin.Context) {
    req := dto.CommentsUpdateReq{}
    s := service.Comments{}
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
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Comments
// @Summary 删除Comments
// @Description 删除Comments
// @Tags Comments
// @Param data body dto.CommentsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/comments [delete]
// @Security Bearer
func (e Comments) Delete(c *gin.Context) {
    s := service.Comments{}
    req := dto.CommentsDeleteReq{}
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
	e.OK( req.GetId(), "删除成功")
}
