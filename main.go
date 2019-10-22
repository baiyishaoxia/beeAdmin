package main

import (
	"beeAdmin/routers"
	_ "beeAdmin/routers"
	"github.com/astaxie/beego"
	"golang.org/x/sync/errgroup"
)
var (
	g errgroup.Group
)
func main() {
	//应用的运行模式，可选值为 prod, dev 或者 test. 默认是 dev, 为开发模式
	if beego.BConfig.RunMode == "dev" {
		//是否开启静态目录的列表显示，默认不显示目录，返回 403 错误。
		beego.BConfig.WebConfig.DirectoryIndex = true
	}
	//当前进程中生成一个home协程
	g.Go(func() error {
		return home()
	})
	beego.Run()
}
//注册web路由
func home() error {
	return routers.InitHomeRouter()
}

