package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Comments2 struct {
	service.Service
}

// GetPage 获取Comments列表
func (e *Comments2) GetPage(c *dto.CommentsGetPageReq2, p *actions.DataPermission, list *[]models.Comments2, count *int64) error {
	var err error

	// 获取帖子ID的全部评论
	err = e.Orm.Model(&models.Comments2{}).Where("blog_id=?", c.BlogID).Where("target_id=?", 0).Find(list).Error
	//err = e.Orm.Model(&data).
	//Scopes(
	//	cDto.MakeCondition(c.GetNeedSearch()),
	//	cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
	//	actions.Permission(data.TableName(), p),
	//).
	//Find(list).Limit(-1).Offset(-1).
	//Count(count).Error
	if err != nil {
		e.Log.Errorf("CommentsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Comments对象
func (e *Comments2) Get(d *dto.CommentsGetReq2, p *actions.DataPermission, model *models.Comments2) error {
	var data models.Comments2

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetComments error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Comments对象
func (e *Comments2) Insert(c *dto.CommentsInsertReq2) error {
	var err error
	var data models.Comments2
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("CommentsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Comments对象
func (e *Comments2) Update(c *dto.CommentsUpdateReq2, p *actions.DataPermission) error {
	var err error
	var data = models.Comments2{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("CommentsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Comments
func (e *Comments2) Remove(d *dto.CommentsDeleteReq2, p *actions.DataPermission) error {
	var data models.Comments2

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveComments error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
