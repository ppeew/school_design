package apis

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"net/http"
	"net/url"
	"os"
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

type User2 struct {
	api.Api
}

// GetPage 获取User列表
// @Summary 获取User列表
// @Description 获取User列表
// @Tags User2
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.User2}} "{"code": 200, "data": [...]}"
// @Router /api/v1/user [get]
// @Security Bearer
func (e User2) GetPage(c *gin.Context) {
	req := dto.UserGetPageReq2{}
	s := service.User2{}
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
	list := make([]models.User2, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取User失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取User
// @Summary 获取User
// @Description 获取User
// @Tags User2
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.User2} "{"code": 200, "data": [...]}"
// @Router /api/v1/user/{id} [get]
// @Security Bearer
func (e User2) Get(c *gin.Context) {
	req := dto.UserGetReq2{}
	s := service.User2{}
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
	var object models.User2

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取User失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建User
// @Summary 创建User
// @Description 创建User
// @Tags User2
// @Accept application/json
// @Product application/json
// @Param data body dto.UserInsertReq2 true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/user [post]
// @Security Bearer
func (e User2) Insert(c *gin.Context) {
	req := dto.UserInsertReq2{}
	s := service.User2{}
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
		e.Error(500, err, fmt.Sprintf("创建User失败，\r\n失败信息 %s", err.Error()))
		return
	}

	//生成token  设置过期时间为1小时
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建JWT的声明
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   gconv.String(req.GetId()),
	}

	// 创建JWT并设置声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名JWT
	tokenString, _ := token.SignedString(jwtKey)

	e.OK(gin.H{
		"id":    req.GetId(),
		"token": tokenString,
	}, "创建成功")
}

var jwtKey = []byte("my_jwt_claims")

// Update 修改User
// @Summary 修改User
// @Description 修改User
// @Tags User2
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.UserUpdateReq2 true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/user/{id} [put]
// @Security Bearer
func (e User2) Update(c *gin.Context) {
	req := dto.UserUpdateReq2{}
	s := service.User2{}
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
		e.Error(500, err, fmt.Sprintf("修改User失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除User
// @Summary 删除User
// @Description 删除User
// @Tags User2
// @Param data body dto.UserDeleteReq2 true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/user [delete]
// @Security Bearer
func (e User2) Delete(c *gin.Context) {
	s := service.User2{}
	req := dto.UserDeleteReq2{}
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
		e.Error(500, err, fmt.Sprintf("删除User失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

func (e User2) Login(c *gin.Context) {
	s := service.User2{}
	req := dto.UserLoginReq2{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Query).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	data, err := s.Login(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("登录，\r\n失败信息 %s", err.Error()))
		return
	}

	//生成token  设置过期时间为1小时
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建JWT的声明
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   gconv.String(data.GetId()),
	}

	// 创建JWT并设置声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名JWT
	tokenString, _ := token.SignedString(jwtKey)

	e.OK(gin.H{
		"result": data,
		"token":  tokenString,
	}, "登录成功")
}

func (e User2) Upload(c *gin.Context) {
	if e.MakeContext(c).MakeOrm().Errors != nil {
		c.String(http.StatusInternalServerError, "orm error")
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	defer file.Close()

	// 创建目标文件
	dir := "/var/image/"
	if config.ApplicationConfig.Mode == "dev" {
		dir = "D:/images/"
	}
	out, err := os.Create(dir + header.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create file")
		return
	}
	defer out.Close()

	// 将上传的文件内容复制到目标文件
	_, err = io.Copy(out, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to save file")
		return
	}

	// 将文件名进行 URL 编码
	filename := url.QueryEscape(header.Filename)

	ip := "139.159.234.134"
	if config.ApplicationConfig.Mode == "dev" {
		ip = "127.0.0.1"
	}
	url := fmt.Sprintf("http://%s:8888/images/%s", ip, filename)

	// 从请求中读取上传的文件
	get := c.Query("type")
	if get == "1" {
		//用户imgae图片上车
		s := c.Query("userId")

		e.Orm.Model(&models.User2{}).Where("id=?", s).Updates(map[string]interface{}{
			"image": url,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}
