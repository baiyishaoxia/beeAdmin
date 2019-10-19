package routers

import (
	_ "beeAdmin/boot"
	"beeAdmin/controllers"
	"beeAdmin/controllers/user"
	"github.com/astaxie/beego"
	"github.com/beego/admin" //admin 包
)

func init() {
	//加载后台权限路由
	admin.Run()
    beego.Router("/home", &controllers.MainController{})

	//加载namespace路由风格
	beego.Router("/admin/user/list", &user.IndexController{},"get:Index")
	//ns := beego.NewNamespace("/admin",
	//	//用户管理
	//	beego.NSRouter("/user/list",&user.IndexController{},"get:Index"),
	//)
	////注册 namespace
	//beego.AddNamespace(ns)
}
