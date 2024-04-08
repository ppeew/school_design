package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/gogf/gf/util/gconv"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Collects struct {
	service.Service
}

// GetPage 获取Collects列表
func (e *Collects) GetPage(c *dto.CollectsGetPageReq, p *actions.DataPermission, list *[]models.Collects, count *int64) error {
	var err error

	tx := e.Orm.Model(&models.Collects{})
	if c.ID != "" {
		tx.Where("user_id=?", gconv.Int(c.ID))
	}
	err = tx.Find(list).Error
	if err != nil {
		e.Log.Errorf("CollectsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Collects对象
func (e *Collects) Get(d *dto.CollectsGetReq, p *actions.DataPermission, model *models.Collects) error {
	var data models.Collects

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetCollects error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Collects对象
func (e *Collects) Insert(c *dto.CollectsInsertReq) error {
	var err error
	var data models.Collects
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("CollectsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Collects对象
func (e *Collects) Update(c *dto.CollectsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Collects{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("CollectsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Collects
func (e *Collects) Remove(d *dto.CollectsDeleteReq, p *actions.DataPermission) error {
	var data models.Collects

	db := e.Orm.Model(&data).Where("user_id=?", d.UserId).Where("blog_id=?", d.BlogId).Delete(&data)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveCollects error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
