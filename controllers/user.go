package controllers

import (
	"book-backstage/models"
	"book-backstage/models/base"
	"github.com/astaxie/beego"
)

// Operations about object
type BookController struct {
	beego.Controller
}


//创建用户信息
//需要传入: 用户名:string  密码:string
func (that *BookController) CreateUser() {
	name := that.GetString("name")
	password := that.GetString("password")

	result := base.Response{}
	if name == "" || password == "" {
		result.Success = false
		result.Msg = "用户名或密码不能为空"
	} else {
		var user models.User
		user.Name= name
		user.Password = password
		// insert
		err := models.InsertUser(user)
		//插入结果
		if err != nil {
			result.Success = false
			result.Msg = err.Error()
		} else {
			result.Success = true
			result.Msg = "用户创建成功"
		}
	}
	that.Data["json"] = result
	that.ServeJSON()
}

//查询用户信息
//输入参数 用户名:name(可选)  不输入用户名,则默认查询所有用户信息
func (that *BookController) QueryUserInfo()  {
	name := that.GetString("name")

	response := base.Response{}
	//不为空定点查询
	if name != ""{
		user, err := models.QueryUserInfo(name)
		if err != nil {
			response.Success = false
			response.Msg = "查询用户信息失败: " + err.Error()
		}else {
			response.Success = true
			response.Msg = "查询成功"
			response.Data = user
		}
	}else {
		//查询所有用户信息
		users, err := models.QueryAllUserInfo()
		if err != nil {
			response.Success = false
			response.Msg = "查询所有用户信息失败: " + err.Error()
		}else {
			response.Success = true
			response.Msg = "查询成功"
			response.Data = users
		}
	}
	that.Data["json"] = response
	that.ServeJSON()
}