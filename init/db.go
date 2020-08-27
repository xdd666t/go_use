package init

import (
	_ "github.com/go-sql-driver/mysql"

	"book_backstage/models"
	"github.com/astaxie/beego/orm"
)

func init() {
	//需要使用的 driver 加入 import 中 mysql:  _ "github.com/go-sql-driver/mysql"
	// 参数1	driverName 	参数2 数据库类型
	// 这个用来设置 driverName 对应的数据库类型
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// ORM 必须注册一个别名为 default 的数据库，作为默认使用。
	// 参数1	数据库的别名，用来在 ORM 中切换数据库使用 	参数2	driverName 	参数3	对应的链接字符串
	_ = orm.RegisterDataBase("default", "mysql",
		"root:246320@tcp(localhost:3306)/book?charset=utf8")
	// 注册模型
	orm.RegisterModel(new(models.User))
	_ = orm.RunSyncdb("default", false, true)
}
