package controllers

import (
	"beeAdmin/common"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//提交视图页面
func (c *MainController) Post() {
	//是否使用XSRF验证
	c.EnableXSRF = true
	//接收数据
	var (
		name  string = c.Input().Get("name")
		token string = c.GetString("_xsrf")
	)
	fmt.Println(name, token)
	data := &common.JSONS{common.HttpSuccess, "获取成功", common.Response{
		List: map[string]string{"name": name, "xsrf": token},
		Time:time.Now().Format("2006-01-02 15:04:05")},
	}
	c.Data["json"] = data
	c.ServeJSON()
}
//测试路由
func (c *MainController) GetTest() {
	item := map[string]string{
		"name":    "apple",
		"version": "10.1",
	}
	data := &common.JSONS{common.HttpSuccess, "获取成功", common.Response{List: item,Time:time.Now().Format("2006-01-02 15:04:05")}}
	c.Data["json"] = data
	c.ServeJSON()
}
