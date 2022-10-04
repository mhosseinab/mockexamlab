package models

type QueryParams struct {
	Search         string `form:"search"`
	Filter         string `form:"filter"`
	Page           int    `form:"page,default=1"`
	PageSize       int    `form:"page_size,default=10"`
	All            bool   `form:"all,default=false"`
	OrderBy        string `form:"order_by,default=id"`
	OrderDirection string `form:"order_direction,default=desc,oneof=desc asc"`
}
