package routers

import (
	. "book-backstage/controllers"
	"github.com/astaxie/beego"
)

func init() {
	router := beego.NewNamespace("/user",
		//创建用户
		beego.NSRouter("/createUser/?:name/?:password",
			&BookController{}, "post:CreateUser",
		),
		//查询用户信息
		beego.NSRouter("/queryUserInfo/?:name",
			&BookController{}, "get:QueryUserInfo",
		),
	)

	beego.AddNamespace(router)
}
