package model

import "time"

// https://www.qetool.com/sql_json_go/sql.html   勾选 json gorm db
// https://www.qetool.com/sql_json_go/sql.html

/**
CREATE TABLE `pf_ticket_info` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_ticket_id` int DEFAULT NULL COMMENT '客户端信息id',
  `station_id` int unsigned NOT NULL COMMENT '站点id',
  `ticket_sn` varchar(32) NOT NULL,
  `lorry_number` varchar(10) NOT NULL,
  `rough_weight` float NOT NULL COMMENT '毛重(千克单位)',
  `tare_weight` float NOT NULL COMMENT '皮重(千克单位)',
  `net_weight` float NOT NULL COMMENT '净重(千克单位)',
  `rough_time` datetime NOT NULL COMMENT '毛重时间',
  `upload_status` int unsigned NOT NULL DEFAULT '0' COMMENT '上报状态，0待上报, 1 重试中  2 上传成功',
  `upload_retry` int unsigned NOT NULL DEFAULT '0' COMMENT '上报尝试次数',
  `photo_front` varchar(255) NOT NULL COMMENT '车前抓拍照片路径 文件名01',
  `photo_behind` varchar(255) NOT NULL COMMENT '车后抓拍照片路径 文件名02',
  `photo_lorry_number` varchar(255) NOT NULL COMMENT '抓车牌拍照片路径 文件名03',
  `photo_goods` varchar(255) NOT NULL COMMENT '货物抓拍照片路径 文件名04',
  `add_time` datetime NOT NULL,
  `update_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `upload_time` timestamp NULL DEFAULT NULL COMMENT '上传成功时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=228 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='信息信息表';
*/
type Ticket struct {
	Id               int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:主键id"`
	StationId        int64     `gorm:"type:int(11);column:station_id;comment:0为正常，非0为禁用时间戳" json:"station_id"`
	ClientTicketId   string    `gorm:"type:int(11);column:client_ticket_id;comment:客户端id" json:"client_ticket_id"`
	TicketSn         string    `gorm:"type:varchar(32);column:ticket_sn;not null;comment:信息号" json:"ticket_sn"`
	LorryNumber      string    `gorm:"type:varchar(32);column:lorry_number;not null;comment:车牌号" json:"lorry_number"`
	RoughWeight      string    `gorm:"type:float(32);column:rough_weight;not null;comment:毛重(千克单位)" json:"rough_weight"`
	TareWeight       string    `gorm:"type:float(32);column:tare_weight;not null;comment:皮重(千克单位)" json:"tare_weight"`
	NetWeight        string    `gorm:"type:float(32);column:net_weight;not null;comment:皮重(千克单位)" json:"net_weight"`
	RoughTime        time.Time `gorm:"type:datetime;column:rough_time;comment:毛重时间" json:"rough_time"`
	UploadStatus     int64     `gorm:"type:int(11);column:upload_status;comment:上报状态，0待上报,1 重试中 2上传成功" json:"upload_status"`
	UploadRetry      int64     `gorm:"type:int(11);column:upload_retry;comment:上报尝试次数" json:"upload_retry"`
	PhotoFront       string    `gorm:"type:varchar(32);column:photo_front;comment:车前抓拍照片路径" json:"photo_front"`
	PhotoBehind      string    `gorm:"type:varchar(32);column:photo_behind;comment:车后抓拍照片路径" json:"photo_behind"`
	PhotoLorryNumber string    `gorm:"type:varchar(32);column:photo_lorry_number;comment:抓车牌拍照片路径" json:"photo_lorry_number"`
	PhotoGoods       string    `gorm:"type:varchar(32);column:photo_goods;comment:货物抓拍照片路径" json:"photo_goods"`
	UploadTime       time.Time `gorm:"type:timestamp;comment:上传成功时间" json:"upload_time,omitempty"`
	AddTime          time.Time `gorm:"type:datetime;comment:创建时间" json:"add_time,omitempty"`
	UpdateTime       time.Time `gorm:"type:datetime;comment:修改时间" json:"update_time,omitempty"`
	Station          Station
}

// TableName 自定义表名
func (Ticket) TableName() string {
	return "pf_ticket_info"
}
