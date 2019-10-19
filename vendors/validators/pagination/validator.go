package pagination

type Pagination struct {
	Page  int `form:"page"`
	Size  int `form:"size"`
	Total int `form:"total"`
}
