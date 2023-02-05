package responseDto

// ResponseConfig  定义返回的数据模型
type ResponseConfig struct {
	UploadAutoRetry     int64 `json:"upload_auto_retry"`
	UploadFailRetryTime int32 `json:"upload_fail_retry_time"`
	HeartbeatTime       int32 `json:"heartbeat_time"`
}
