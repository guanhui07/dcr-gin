package responseDto

type ResponsePageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageRows int         `json:"page_rows"`
}
