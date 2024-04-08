package service

import (
	"errors"
	"github.com/gogf/gf/util/gconv"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Blogs struct {
	service.Service
}

// GetPage 获取Blogs列表
func (e *Blogs) GetPage(c *dto.BlogsGetPageReq, p *actions.DataPermission, list *[]models.Blogs, count *int64) error {
	var err error
	var data models.Blogs

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("created_at desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BlogsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Blogs对象
func (e *Blogs) Get(d *dto.BlogsGetReq, p *actions.DataPermission, model *models.Blogs) error {
	var data models.Blogs

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBlogs error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Blogs对象
func (e *Blogs) Insert(c *dto.BlogsInsertReq) error {
	var err error
	var data models.Blogs
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BlogsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Blogs对象
func (e *Blogs) Update(c *dto.BlogsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Blogs{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	if c.Type == "1" {
		//点赞操作
		e.Orm.Model(&models.Blogs{}).Where("id=?", c.Id).Update("star", gconv.Int(data.Star)+1)

	} else {
		db := e.Orm.Save(&data)
		if err = db.Error; err != nil {
			e.Log.Errorf("BlogsService Save error:%s \r\n", err)
			return err
		}
		if db.RowsAffected == 0 {
			return errors.New("无权更新该数据")
		}
		return nil
	}
	return nil
}

// Remove 删除Blogs
func (e *Blogs) Remove(d *dto.BlogsDeleteReq, p *actions.DataPermission) error {
	var data models.Blogs

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveBlogs error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
