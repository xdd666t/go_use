package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type User struct {
	Id       int `orm:"pk;auto"` // 设置为主键，字段Id
	Name string
	Password string
}

//向数据中插入新用户
func InsertUser(user User) error {
	//判断用户名是否重复
	repeatUser, _ := QueryUserInfo(user.Name)
	if  repeatUser != (User{}) {
		return errors.New("用户名重复，请更换用户名后重试")
	}
	//插入用户
	o := orm.NewOrm()
	// insert
	_, err := o.Insert(&user)
	return err
}

 //通过用户名获取用户信息
func QueryUserInfo(name string) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("name", name).One(&user)
	return user, err
}

//获取全部用户信息
func QueryAllUserInfo() ([]User, error) {
	var user []User
	o := orm.NewOrm()
	_, err := o.QueryTable("user").All(&user)
	return user, err
}