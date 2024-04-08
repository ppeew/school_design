package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/models"
)

type Message struct {
	api.Api
}

func (m Message) GetPage(c *gin.Context) {
	m.MakeContext(c).MakeOrm()

	req := struct {
		UserID int `form:"userId"`
	}{}

	c.BindQuery(&req)

	ret := make([]models.Message, 0)
	m.Orm.Model(&models.Message{}).Where("user_id=?", req.UserID).Find(&ret)
	var count int64
	m.Orm.Model(&models.Message{}).Where("user_id=?", req.UserID).Where("notice_status=?", 0).Count(&count)

	m.OK(gin.H{
		"list":        ret,
		"unReadCount": count,
	}, "ok")
}

func (m Message) Get(c *gin.Context) {
	m.MakeContext(c).MakeOrm()

	req := struct {
		UserID int `form:"userId"`
	}{}

	c.BindQuery(&req)

	ret := make([]models.Message, 0)
	m.Orm.Model(&models.Message{}).Where("user_id=?", req.UserID).Find(&ret)

	m.OK(ret, "ok")
}

func (m Message) Update(c *gin.Context) {
	m.MakeContext(c).MakeOrm()

	req := struct {
		MessageID int `uri:"id"`
	}{}

	c.BindUri(&req)

	m.Orm.Model(&models.Message{}).Where("id=?", req.MessageID).Update("notice_status", 1)

	m.OK(req.MessageID, "ok")
}

func (m Message) Delete(c *gin.Context) {
	m.MakeContext(c).MakeOrm()

	value := c.Query("id")

	m.Orm.Model(&models.Message{}).Where("id=?", value).Delete(&models.Message{})

	m.OK(value, "ok")
}

func (m Message) AllRead(c *gin.Context) {
	m.MakeContext(c).MakeOrm()

	value := c.Query("userId")
	m.Orm.Model(&models.Message{}).Where("user_id=?", value).Update("notice_status", 1)

	m.OK(value, "ok")
}

func (m Message) DeleteAll(c *gin.Context) {
	m.MakeContext(c).MakeOrm()

	value := c.Query("userId")

	m.Orm.Model(&models.Message{}).Where("user_id=?", value).Delete(&models.Message{})

	m.OK(value, "ok")
}
