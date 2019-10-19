package user

import "beeAdmin/vendors/validators/pagination"

type GetIndexParam struct {
	Lang string `form:"lang" validate:"required"` //必填项
	Keywords string `form:"keywords" validate:"omitempty,min=1,max=15"` //非必填，若有值则进行后续验证
	pagination.Pagination
}
