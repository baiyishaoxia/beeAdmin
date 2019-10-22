package routers

import (
	"beeAdmin/controllers"
	"beeAdmin/controllers/user"
	"github.com/astaxie/beego"
)

func InitHomeRouter() error {
	//Session初始化
	beego.BConfig.WebConfig.Session.SessionOn = true
	//前台请求日志
	beego.SetLogger(beego.AppConfig.String("logsAdapterFile"), `{"filename":"runtime/logs/test.log"}`)
	//静态资源
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/uploads", "uploads")

	//加载namespace路由风格
	ns := beego.NewNamespace("/home",
		beego.NSRouter("/", &controllers.MainController{}, "get:Get;post:Post"),
	)
	//注册 namespace
	beego.AddNamespace(ns)

	//加载api路由
	api := beego.NewNamespace("/v1",
		beego.NSRouter("/test/:key([0-9]+)", &controllers.MainController{}, "get:GetTest"),  //测试
	)
	beego.AddNamespace(api)

	//加载普通路由
	beego.Router("/user/index", &user.IndexController{}, "get:Index") //用户列表

	return nil
}
