package requestDto

type StationTicketReq struct {
	StationId string `json:"station_id" binding:"required,min=1,max=11"`
	TicketId  string `json:"ticket_id" binding:"required,min=1,max=11"`
}

/**
{
"station_id":"fsd",
"ticket_id":"sdf",
}
*/

type TicketUploadReq struct {
	Device   string `json:"device" binding:"required,min=1,max=11"`
	TicketId string `json:"ticket_id" binding:"required,min=1,max=11"`
}

/**
{
"device":"fsd",
"ticket_id":"12",
}
*/

type TicketApiReq struct {
	Device      string `json:"device" binding:"required,min=1,max=11"`
	TicketId    string `json:"ticket_id" binding:"required,min=1,max=11"`
	TicketSn    string `json:"ticket_sn" binding:"required,min=1,max=32"`
	LorryNumber string `json:"number_plate" binding:"required,min=1,max=11"`
	RoughWeight string `json:"rough_weight"`
	TareWeight  string `json:"tare_weight"`
	NetWeight   string `json:"net_weight"`
	Time        string `json:"time" `
}
type TicketListReq struct {
	PageInfo
	StationId    int64   `json:"station_id" form:"station_id"`
	TicketSn     string  `json:"ticket_sn" form:"ticket_sn"`
	LorryNumber  string  `json:"lorry_number" form:"lorry_number"`
	NetWeightS   float64 `json:"net_weight_s" form:"net_weight_s" `
	NetWeightE   float64 `json:"net_weight_e" form:"net_weight_e" `
	RoughWeightS float64 `json:"rough_weight_s" form:"rough_weight_s" `
	RoughWeightE float64 `json:"rough_weight_e" form:"rough_weight_e" `
	UploadStatus float64 `json:"upload_status" form:"upload_status" `
	RoughTimeS   string  `json:"rough_time_s" form:"rough_time_s" `
	RoughTimeE   string  `json:"rough_time_e" form:"rough_time_e" `
}
