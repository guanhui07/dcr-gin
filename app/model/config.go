package model

// TableName 自定义表名
func (Config) TableName() string {
	return "pf_config"
}

// https://www.qetool.com/sql_json_go/sql.html   勾选 json gorm db
// http://sql2struct.atotoa.com/
// http://sql2struct.atotoa.com/

/**
CREATE TABLE `pf_config` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `upload_auto_retry` int unsigned NOT NULL DEFAULT '0' COMMENT '上报失败自动重试：0启用，非0为禁用的用户id',
  `upload_fail_retry_time` smallint unsigned NOT NULL DEFAULT '5' COMMENT '上报失败重试间隔，单位秒',
  `heartbeat_time` smallint DEFAULT '30' COMMENT '心跳间隔，单位：秒',
  `add_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统设置表';
*/
// Config 配置表
type Config struct {
	Id                  int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:主键id"`
	UploadAutoRetry     int64  `json:"upload_auto_retry" gorm:"comment:上报失败自动重试：0启用，非0为禁用的用户id"`
	UploadFailRetryTime int32  `json:"upload_fail_retry_time" gorm:"comment:上报失败重试间隔，单位秒"`
	HeartbeatTime       int32  `json:"heartbeat_time" gorm:"comment:心跳间隔，单位：秒"`
	AddTime             string `json:"add_time" gorm:"comment:创建时间;"`
	UpdateTime          string `json:"update_time" gorm:"comment:更新时间"`
}
