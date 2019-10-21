package main

import (
	_ "beeAdmin/routers"
	"github.com/astaxie/beego"
)

func main() {
	//应用的运行模式，可选值为 prod, dev 或者 test. 默认是 dev, 为开发模式
	if beego.BConfig.RunMode == "dev" {
		//是否开启静态目录的列表显示，默认不显示目录，返回 403 错误。
		beego.BConfig.WebConfig.DirectoryIndex = true
	}
	beego.Run()
}

