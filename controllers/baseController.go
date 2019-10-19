package controllers

import (
	"beeAdmin/boot"
	"beeAdmin/common"
	"beeAdmin/vendors/validators"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"gopkg.in/go-playground/validator.v9"
)

//基础controller，用于API继承
type BaseController struct {
	beego.Controller
	i18n.Locale
	common.Responser
}

//在执行方法前
type IBeforeAction interface {
	BeforeAction()
}

//在执行方法后
type IAfterAction interface {
	AfterAction()
}

//这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些Method方法之前执行，用户重写这个函数实现类似用户验证之类。
func (base *BaseController) Prepare() {
	base.setLangVer()
	//使用默认布局页面
	base.Layout = "layouts/main.html"
	if app, ok := base.AppController.(IBeforeAction); ok {
		app.BeforeAction()
	}
}

//这个函数是在执行完相应的HTTP Method方法之后执行的，默认是空，用户可以在子struct中重写这个函数，例如关闭数据库、清理数据之类的工作。
func (base *BaseController) Finish() {
	if app, ok := base.AppController.(IAfterAction); ok {
		app.AfterAction()
	}
}

func (base *BaseController) setLangVer() bool {
	currLocale := boot.GetLocale()
	base.Lang = currLocale.Lang
	base.Responser.Lang = currLocale.Lang
	base.Data["Lang"] = currLocale.Lang
	base.Data["CurrLangName"] = currLocale.Name
	base.Data["RestLocale"] = boot.GetRestLocale()
	return currLocale.IsNeedRedirect
}

//返回json
func (base *BaseController) returnJson() {
	base.ServeJSON()
	base.StopRun()
}

//成功返回
func (base *BaseController) SuccessJSON(status int, msg string, data interface{}) {
	base.Data["json"] = base.Success(status, msg, data)
	base.returnJson()
}

//失败返回
func (base *BaseController) ErrorJSON(status int, msg string, tr ...bool) {
	if beego.BConfig.RunMode != beego.DEV && status == common.QUERY_ERROR {
		status = common.SYSTEM_ERROR
		msg = "system error"
		tr[0] = true
	}
	translate := len(tr) > 0 && tr[0]
	base.Data["json"] = base.Error(status, msg, translate)
	base.returnJson()
}

func (base *BaseController) InvalidArgumentJSON(errors ...string) {
	base.Data["json"] = base.InvalidArgument(errors...)
	base.returnJson()
}

func (base *BaseController) SystemErrorJSON(errors ...string) {
	base.Data["json"] = base.SystemError(errors...)
	base.returnJson()
}

func (base *BaseController) QueryErrorJSON(errors ...string) {
	if beego.BConfig.RunMode != beego.DEV {
		errors = []string{}
	}
	base.Data["json"] = base.QueryError(errors...)
	base.returnJson()
}

func (base *BaseController) Valid(obj interface{}) {
	valid := boot.GetValidator()
	if err := valid.Struct(obj); err != nil {
		errs := err.(validator.ValidationErrors)
		trans := validators.GetTrans(boot.GetLang())
		errTrs := errs.Translate(trans)
		var errors []string
		for _, v := range errTrs {
			errors = append(errors, v)
		}
		base.InvalidArgumentJSON(errors...)
	}
}

func (base *BaseController) ValidForm(obj interface{}) {
	if err := base.ParseForm(obj); err != nil {
		if beego.BConfig.RunMode == beego.DEV {
			base.InvalidArgumentJSON(err.Error())
		}
		base.InvalidArgumentJSON()
	}
	base.Valid(obj)
}

func (base *BaseController) ValidJSON(obj interface{}) {
	if err := json.Unmarshal(base.Ctx.Input.RequestBody, obj); err != nil {
		if beego.BConfig.RunMode == beego.DEV {
			base.InvalidArgumentJSON(err.Error())
		}
		base.InvalidArgumentJSON()
	}
	base.Valid(obj)
}