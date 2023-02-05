package requestDto

type StationSign struct {
	Device string `json:"device" binding:"required"`
	Sign   string `json:"sign" binding:"required"`
}

/**
{
"device":"fsd",
"sign":"sdf",
}
*/
