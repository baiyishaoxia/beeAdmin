package lang

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"
	"github.com/beego/i18n"
	"strings"
)

var (
	AllLang  []Locale
	RestLang []Locale
	Lang     Locale
)

type Locale struct {
	Lang           string
	Name           string
	IsNeedRedirect bool
}

func RegisterLangConf() {
	var appendFiles []string
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	fileTypes := strings.Split(beego.AppConfig.String("lang::appendFileTypes"), "|")
	AllLang = make([]Locale, 0, len(langs))
	for i, v := range langs {
		AllLang = append(AllLang, Locale{
			Lang: v,
			Name: names[i],
		})
	}

	for _, langType := range AllLang {
		lang := langType.Lang
		beego.Trace("Loading language: " + lang)
		localePath := GetLangPath(lang, "locale")
		for _, fileType := range fileTypes {
			appendFile := GetLangPath(lang, fileType)
			if !utils.FileExists(appendFile) {
				continue
			}
			appendFiles = append(appendFiles, GetLangPath(lang, fileType))
		}
		if err := i18n.SetMessage(lang, localePath, appendFiles...); err != nil {
			beego.Error(err.Error())
			return
		}
	}
}

func InitLang(ctx *context.Context) (lang Locale, restLang []Locale) {
	isNeedRedirect := false
	hasCookie := false

	langStr := ctx.Input.Query("lang")
	if len(langStr) == 0 {
		langStr = ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedirect = true
	}

	if !i18n.IsExist(langStr) {
		langStr = ""
		isNeedRedirect = false
		hasCookie = false
	}

	if len(langStr) == 0 {
		al := ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				langStr = al
			}
		}
	}

	if len(langStr) == 0 {
		langStr = "en-US"
		isNeedRedirect = false
	}

	Lang = Locale{
		Lang: langStr,
	}

	if !hasCookie {
		expired := 7 * 24 * 3600
		ctx.SetCookie("lang", Lang.Lang, expired, "/")
	}

	RestLang = make([]Locale, 0, len(AllLang)-1)
	for _, v := range AllLang {
		if langStr != v.Lang {
			RestLang = append(RestLang, v)
		} else {
			Lang.Name = v.Name
		}
	}

	Lang.IsNeedRedirect = isNeedRedirect

	return Lang, RestLang
}

func GetLangPath(lang, kind string) string {
	return fmt.Sprintf("conf/languages/%s/%s.ini", lang, kind)
}
