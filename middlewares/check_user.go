package middlewares

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

//region   登录前验证   Author:tang
func CheckLogin() func(ctx *context.Context){
	return func(ctx *context.Context) {
		fmt.Println("在这里可以做一些AdminLogin处理")
	}
}
//endregion