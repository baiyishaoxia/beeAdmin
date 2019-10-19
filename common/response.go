package common

import (
	"beeAdmin/common/helper"
	"github.com/beego/i18n"
)

//响应体
type Responser struct {
	JSONS
	i18n.Locale
}

//状态码200、提示信息，数据
func (r *Responser) Success(status int, msg string, data interface{}) JSONS {
	r.Status = status
	r.Msg = r.Tr(msg)
	r.Data.List = data
	r.Data.Time = helper.Date("Y-m-d H:i:s")
	return r.JSONS
}

//状态码201，提示信息
func (r *Responser) Error(err int, message string, tr ...bool) JSONS {
	r.Status = err
	if len(tr) > 0 && tr[0] {
		r.Msg = r.Tr(message)
	} else {
		r.Msg = message
	}
	r.Data.Time = helper.Date("Y-m-d H:i:s")
	r.Data.List = new(struct{})
	return r.JSONS
}

//抛出异常
func (res *JSONS) parseErrors(errors ...string) {
	if len(errors) > 0 {
		res.Data.Errors = errors
	}
}

func (r *Responser) InvalidArgument(errors ...string) JSONS {
	res := r.Error(PARAMS_ERROR, "params error", true)
	res.parseErrors(errors...)
	return res
}

func (r *Responser) SystemError(errors ...string) JSONS {
	res := r.Error(SYSTEM_ERROR, "system error", true)
	res.parseErrors(errors...)
	return res
}

func (r *Responser) QueryError(errors ...string) JSONS {
	res := r.Error(QUERY_ERROR, "query error", true)
	res.parseErrors(errors...)
	return res
}
