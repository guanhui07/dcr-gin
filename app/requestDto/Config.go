package requestDto

type CreateConfigReq struct {
	UploadAutoRetry     int64 `json:"upload_auto_retry" binding:"gte=0"`      // 上报失败自动重试：0启用，非0为禁用的用户id
	UploadFailRetryTime int32 `json:"upload_fail_retry_time" binding:"gte=0"` //上报失败重试间隔，单位秒
	HeartbeatTime       int32 `json:"heartbeat_time" binding:"gte=0"`         //心跳间隔，单位：秒
}

/**
{
"upload_auto_retry":2,
"upload_fail_retry_time":2,
"heartbeat_time":2
}
*/
