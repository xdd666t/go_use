package routers

import (
	"book-backstage/controllers"
	"github.com/astaxie/beego"
)

func init() {
	router := beego.NewNamespace("/user",
		//创建用户
		beego.NSRouter("/addUser/?:name/?:password",
			&controllers.BookController{}, "post:AddUser",
		),
		//查询用户信息
		beego.NSRouter("/getUserInfo/?:name",
			&controllers.BookController{}, "get:GetUserInfo",
		),
	)

	beego.AddNamespace(router)
}
