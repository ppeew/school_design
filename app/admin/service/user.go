package service

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
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

	//先查询用户是否存在
	var u models.User
	result := e.Orm.Where("username = ?", c.Username).First(&u)
	if result.RowsAffected == 1 {
		return errors.New("重复创建用户")
	}
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encoded := password.Encode(c.Password, options)
	encodePassword := fmt.Sprintf("%s$%s", salt, encoded)
	data.Password = encodePassword

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

	if c.Password != "" {
		options := &password.Options{
			SaltLen:      16,
			Iterations:   100,
			KeyLen:       32,
			HashFunction: sha512.New,
		}
		salt, encoded := password.Encode(c.Password, options)
		encodePassword := fmt.Sprintf("%s$%s", salt, encoded)
		data.Password = encodePassword
	}

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

func (e *User) Login(d *dto.UserLoginReq) (*models.User, error) {
	// 获得该用户的username，password
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}

	data := new(models.User)
	if tx := e.Orm.Model(models.User{}).Where("username=?", d.Username).First(data); tx.Error != nil {
		return nil, tx.Error
	}

	if data.Id == 0 {
		return nil, errors.New("无该用户")
	}

	info := strings.Split(data.Password, "$")
	verify := password.Verify(d.PassWord, info[0], info[1], options)

	if verify {
		return data, nil
	}
	return nil, errors.New("密码错误")
}
