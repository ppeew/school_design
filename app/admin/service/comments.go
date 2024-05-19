package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Comments struct {
	service.Service
}

// GetPage 获取Comments列表
func (e *Comments) GetPage(c *dto.CommentsGetPageReq, p *actions.DataPermission, list *[]models.Comments, count *int64) error {
	var err error
	var data models.Comments

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("CommentsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Comments对象
func (e *Comments) Get(d *dto.CommentsGetReq, p *actions.DataPermission, model *models.Comments) error {
	var data models.Comments

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
func (e *Comments) Insert(c *dto.CommentsInsertReq) error {
    var err error
    var data models.Comments
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("CommentsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Comments对象
func (e *Comments) Update(c *dto.CommentsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Comments{}
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
func (e *Comments) Remove(d *dto.CommentsDeleteReq, p *actions.DataPermission) error {
	var data models.Comments

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
