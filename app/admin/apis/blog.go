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

type Blogs struct {
	api.Api
}

// GetPage 获取Blogs列表
// @Summary 获取Blogs列表
// @Description 获取Blogs列表
// @Tags Blogs
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Blogs}} "{"code": 200, "data": [...]}"
// @Router /api/v1/blogs [get]
// @Security Bearer
func (e Blogs) GetPage(c *gin.Context) {
	req := dto.BlogsGetPageReq{}
	s := service.Blogs{}
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
	list := make([]models.Blogs, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Blogs失败，\r\n失败信息 %s", err.Error()))
		return
	}

	for _, blogs := range list {
		blogs.CreatedAt.Format(time.DateTime)
	}

	type resp struct {
		ID         int
		Username   string
		Comment    int
		Star       int
		Collect    int
		CreateTime string
		AvatarUrl  string
		Title      string
	}

	ret := make([]resp, 0)
	for _, blogs := range list {
		var commentCount int64
		e.Orm.Model(models.Comments{}).Where("blog_id=?", blogs.Id).Count(&commentCount)

		var collectCount int64
		e.Orm.Model(models.Collects{}).Where("blog_id=?", blogs.Id).Count(&collectCount)

		// 查询该username对应的id
		us := new(models.User)
		e.Orm.Where("username=?", blogs.Username).First(us)

		re := resp{
			ID:         blogs.Id,
			Username:   blogs.Username,
			Comment:    int(commentCount),
			Star:       gconv.Int(blogs.Star),
			Collect:    int(collectCount),
			CreateTime: blogs.CreatedAt.Format(time.DateTime),
			AvatarUrl:  us.Image,
			Title:      blogs.Title,
		}
		if us.Image == "" {
			re.AvatarUrl = "https://th.bing.com/th/id/R.ec214a16a4b823260966fdb7b09eddff?rik=ivXWTf6iBDSAsg&riu=http%3a%2f%2fpic2.nipic.com%2f20090429%2f984755_194557035_2.jpg&ehk=u7%2bUd79GB72em%2b7h%2bdvk6exyLNUKzpSdUpeA8DWRn2s%3d&risl=&pid=ImgRaw&r=0"
		}

		ret = append(ret, re)
	}

	e.PageOK(ret, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Blogs
// @Summary 获取Blogs
// @Description 获取Blogs
// @Tags Blogs
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Blogs} "{"code": 200, "data": [...]}"
// @Router /api/v1/blogs/{id} [get]
// @Security Bearer
func (e Blogs) Get(c *gin.Context) {
	req := dto.BlogsGetReq{}
	s := service.Blogs{}
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
	var object models.Blogs

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Blogs失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Blogs
// @Summary 创建Blogs
// @Description 创建Blogs
// @Tags Blogs
// @Accept application/json
// @Product application/json
// @Param data body dto.BlogsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/blogs [post]
// @Security Bearer
func (e Blogs) Insert(c *gin.Context) {
	req := dto.BlogsInsertReq{}
	s := service.Blogs{}
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
		e.Error(500, err, fmt.Sprintf("创建Blogs失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Blogs
// @Summary 修改Blogs
// @Description 修改Blogs
// @Tags Blogs
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BlogsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/blogs/{id} [put]
// @Security Bearer
func (e Blogs) Update(c *gin.Context) {
	req := dto.BlogsUpdateReq{}
	s := service.Blogs{}
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
		e.Error(500, err, fmt.Sprintf("修改Blogs失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Blogs
// @Summary 删除Blogs
// @Description 删除Blogs
// @Tags Blogs
// @Param data body dto.BlogsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/blogs [delete]
// @Security Bearer
func (e Blogs) Delete(c *gin.Context) {
	s := service.Blogs{}
	req := dto.BlogsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Blogs失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
