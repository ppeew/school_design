package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type UserGetPageReq struct {
	dto.Pagination `search:"-"`
	UserOrder
}

type UserOrder struct {
	Id       string `form:"idOrder"  search:"type:order;column:id;table:user"`
	Username string `form:"usernameOrder"  search:"type:order;column:username;table:user"`
	Password string `form:"passwordOrder"  search:"type:order;column:password;table:user"`
	Age      string `form:"ageOrder"  search:"type:order;column:age;table:user"`
	Sex      string `form:"sexOrder"  search:"type:order;column:sex;table:user"`
}

func (m *UserGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type UserInsertReq struct {
	Id       int    `json:"-" comment:""` //
	Username string `json:"username" comment:"用户名"`
	Password string `json:"password" comment:"密码"`
	Age      string `json:"age" comment:"年龄"`
	Sex      string `json:"sex" comment:"性别,1:男 2:女"`
	common.ControlBy
}

func (s *UserInsertReq) Generate(model *models.User) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Password = s.Password
	model.Age = s.Age
	model.Sex = s.Sex
}

func (s *UserInsertReq) GetId() interface{} {
	return s.Id
}

type UserUpdateReq struct {
	Id       int    `uri:"id" comment:""` //
	Username string `json:"username" comment:"用户名"`
	Password string `json:"password" comment:"密码"`
	Age      string `json:"age" comment:"年龄"`
	Sex      string `json:"sex" comment:"性别,1:男 2:女"`
	common.ControlBy
}

func (s *UserUpdateReq) Generate(model *models.User) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	//model.Password = s.Password
	model.Age = s.Age
	model.Sex = s.Sex
}

func (s *UserUpdateReq) GetId() interface{} {
	return s.Id
}

// UserGetReq 功能获取请求参数
type UserGetReq struct {
	Id int `uri:"id"`
}

func (s *UserGetReq) GetId() interface{} {
	return s.Id
}

// UserDeleteReq 功能删除请求参数
type UserDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *UserDeleteReq) GetId() interface{} {
	return s.Ids
}

type UserLoginReq struct {
	Username string `form:"username" json:"username"`
	PassWord string `form:"password" json:"password"`
}

func (u UserLoginReq) Check() bool {
	return true
}
