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

type User struct {
	service.Service
}

// GetPage 获取User列表
func (e *User) GetPage(c *dto.UserGetPageReq, p *actions.DataPermission, list *[]models.User, count *int64) error {
	var err error
	var data models.User

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("UserService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取User对象
func (e *User) Get(d *dto.UserGetReq, p *actions.DataPermission, model *models.User) error {
	var data models.User

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetUser error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建User对象
func (e *User) Insert(c *dto.UserInsertReq) error {
    var err error
    var data models.User
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("UserService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改User对象
func (e *User) Update(c *dto.UserUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.User{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("UserService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除User
func (e *User) Remove(d *dto.UserDeleteReq, p *actions.DataPermission) error {
	var data models.User

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveUser error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
