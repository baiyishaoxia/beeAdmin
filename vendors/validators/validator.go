package validators

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	enTranslation "gopkg.in/go-playground/validator.v9/translations/en"
	jaTranslation "gopkg.in/go-playground/validator.v9/translations/ja"
	zhTranslation "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"strings"
)

var (
	validate *validator.Validate
	enTrans  ut.Translator
	zhTrans  ut.Translator
	jaTrans  ut.Translator
)

func InitValidate() *validator.Validate {
	validate = validator.New()

	// 验证器字段根据标签 form 或者 json 的名字定
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		getFieldName := func(str string) string {
			return strings.SplitN(fld.Tag.Get(str), ",", 2)[0]
		}
		name := getFieldName("json")
		if name == "" {
			name = getFieldName("form")
		}
		if name == "-" {
			return ""
		}
		return name
	})

	initUniversalValidator()

	return validate
}

func GetTrans(lang string) ut.Translator {
	switch lang {
	case "en-US":
		return enTrans
	case "zh-CN":
		return zhTrans
	case "ja":
		return jaTrans
	default:
		return enTrans
	}
}

func initUniversalValidator() {
	enTrans = getEnTrans()
	zhTrans = getZhTrans()
	jaTrans = getJaTrans()
	_ = enTranslation.RegisterDefaultTranslations(validate, enTrans)
	_ = zhTranslation.RegisterDefaultTranslations(validate, zhTrans)
	_ = jaTranslation.RegisterDefaultTranslations(validate, jaTrans)
}

func getEnTrans() ut.Translator {
	var uni *ut.UniversalTranslator
	enTr := en.New()
	uni = ut.New(enTr, enTr)
	trans, _ := uni.GetTranslator("en")
	return trans
}

func getZhTrans() ut.Translator {
	var uni *ut.UniversalTranslator
	zhTr := zh.New()
	uni = ut.New(zhTr, zhTr)
	trans, _ := uni.GetTranslator("zh")
	return trans
}

func getJaTrans() ut.Translator {
	var uni *ut.UniversalTranslator
	jaTr := ja.New()
	uni = ut.New(jaTr, jaTr)
	trans, _ := uni.GetTranslator("ja")
	return trans
}
