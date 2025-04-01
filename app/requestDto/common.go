package requestDto

// https://oktools.net/json2go   json转结构体

// PageInfo common
type PageInfo struct {
	Page     int    `json:"page,default=1" form:"page,default=1"` // 页码
	PageRows int    `json:"page_rows,default=10" form:"page_rows,default=10" `
	Keyword  string `json:"keyword" form:"keyword"` //关键字
}

/**
{
"page",1,
"page_rows",10,
"keyword",10,
}
*/

// GetByIdReq Find by id structure
type GetByIdReq struct {
	Id int `json:"id" form:"id"` // 主键Id
}

/**
{
"id":1,
}
*/

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
