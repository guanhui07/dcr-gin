package requestDto

type MyJson struct {
	v any
}

type AddStationReq struct {
	Name string `json:"name" binding:"required"` //  站点名称
	//Location    string `json:"location" binding:"required"` // 坐标
	IpAddress   string `json:"ip_address" binding:"required"`          // 站点ip地址,非此IP数据不接收
	TigerShaped string `json:"tiger_shaped" binding:"required,len=32"` // 握手符号 32位
	//Heartbeat   int64 `json:"heartbeat" binding:"gte=0"` // 上次心跳时间戳
	Lng float32 `json:"lng" binding:"min=0,max=180" `
	Lat float32 `json:"lat" binding:"min=0,max=180"`
	//Status 		int64 `json:"status" binding:"required,len=10"` //站点状态：0正常，非0为停用时间戳
}

/**
{
"name":"fsd",
"ip_address":"sdf",
"tiger_shaped":"sdf",
"lng":"sdf",
"lat":"sdf",
}
*/

type EditStationReq struct {
	Name        string  `json:"name"`         //  站点名称
	IpAddress   string  `json:"ip_address"`   // 站点ip地址,非此IP数据不接收
	TigerShaped string  `json:"tiger_shaped"` // 握手符号 32位
	Lng         float32 `json:"lng" binding:"min=0,max=180" `
	Lat         float32 `json:"lat" binding:"min=0,max=180"`
	Id          int64   `json:"id" binding:"required,gt=0"`
}

/**
{
"name":"fsd",
"ip_address":"sdf",
"tiger_shaped":"sdf",
"lng":"sdf",
"lat":"sdf",
"id":"sdf",
}
*/

type StationListReq struct {
	Page     int `json:"page,default=1" form:"page,default=1" `
	PageRows int `json:"page_rows,default=10" form:"page_rows,default=10" `
}

type StationStatusReq struct {
	Id     int64 `json:"id" binding:"required,gt=0"`
	Status int64 `json:"status" binding:"gte=0"`
}
