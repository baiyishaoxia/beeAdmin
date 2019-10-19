package search

import "beeAdmin/vendors/validators/pagination"

type SearchParam struct {
	Word string `form:"word" valid:"required"`
	Lang string `form:"lang" valid:"required"`
	pagination.Pagination
}