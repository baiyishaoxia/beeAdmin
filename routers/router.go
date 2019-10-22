package routers

import (
	"beeAdmin/boot"
	_ "beeAdmin/boot"
	"beeAdmin/controllers"
	"beeAdmin/controllers/user"
	"beeAdmin/middlewares"
	"beeAdmin/vendors/geoip"
	"beeAdmin/vendors/lang"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/beego/admin" //admin 包
)

func init() {
	//加载后台权限路由
	admin.Run()
    beego.Router("/home", &controllers.MainController{})
	//beego.Router("/admin/user/list", &user.IndexController{},"get:Index")

	beego.InsertFilter("/*", beego.BeforeRouter, func(context *context.Context) {
		// 获取当前语言
		boot.App.Locale, boot.App.RestLocale = lang.InitLang(context)
		boot.App.GeoIP = geoip.InitGeoIP(context)
	})
	//加载namespace路由风格
	ns := beego.NewNamespace("/admin",
		//用户管理模块
		beego.NSRouter("/user/list",&user.IndexController{},"get:Index"),
	)
	//注册 namespace
	beego.AddNamespace(ns)

	//中间件控制
	beego.InsertFilter("/admin/user/list",beego.BeforeRouter,middlewares.CheckLogin())
}
