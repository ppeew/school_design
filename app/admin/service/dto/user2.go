package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type UserGetPageReq2 struct {
	dto.Pagination `search:"-"`
	UserOrder2
}

type UserOrder2 struct {
	Id       string `form:"idOrder"  search:"type:order;column:id;table:user"`
	Username string `form:"usernameOrder"  search:"type:order;column:username;table:user"`
	Password string `form:"passwordOrder"  search:"type:order;column:password;table:user"`
	Age      string `form:"ageOrder"  search:"type:order;column:age;table:user"`
	Sex      string `form:"sexOrder"  search:"type:order;column:sex;table:user"`
}

func (m *UserGetPageReq2) GetNeedSearch() interface{} {
	return *m
}

type UserInsertReq2 struct {
	Id       int    `json:"-" comment:""` //
	Username string `json:"username" comment:"用户名"`
	Password string `json:"password" comment:"密码"`
	Age      string `json:"age" comment:"年龄"`
	Sex      string `json:"sex" comment:"性别,1:男 2:女"`
	common.ControlBy
}

func (s *UserInsertReq2) Generate(model *models.User2) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Password = s.Password
	model.Age = s.Age
	model.Sex = s.Sex
}

func (s *UserInsertReq2) GetId() interface{} {
	return s.Id
}

type UserUpdateReq2 struct {
	Id       int    `uri:"id" comment:""` //
	Username string `json:"username" comment:"用户名"`
	Password string `json:"password" comment:"密码"`
	Age      string `json:"age" comment:"年龄"`
	Sex      string `json:"sex" comment:"性别,1:男 2:女"`
	common.ControlBy
}

func (s *UserUpdateReq2) Generate(model *models.User2) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	//model.Password = s.Password
	model.Age = s.Age
	model.Sex = s.Sex
}

func (s *UserUpdateReq2) GetId() interface{} {
	return s.Id
}

// UserGetReq2 功能获取请求参数
type UserGetReq2 struct {
	Id int `uri:"id"`
}

func (s *UserGetReq2) GetId() interface{} {
	return s.Id
}

// UserDeleteReq2 功能删除请求参数
type UserDeleteReq2 struct {
	Ids []int `json:"ids"`
}

func (s *UserDeleteReq2) GetId() interface{} {
	return s.Ids
}

type UserLoginReq2 struct {
	Username string `form:"username" json:"username"`
	PassWord string `form:"password" json:"password"`
}

func (u UserLoginReq2) Check() bool {
	return true
}
