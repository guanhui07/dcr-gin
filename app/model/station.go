package model

import (
	"gorm.io/datatypes"
	"time"
)

// https://www.qetool.com/sql_json_go/sql.html   勾选 json gorm db
// https://www.qetool.com/sql_json_go/sql.html

/**
CREATE TABLE `pf_station_info` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '站点名称',
  `location` json NOT NULL COMMENT '站点所在经纬度',
  `ip_address` varchar(15) NOT NULL DEFAULT '' COMMENT '站点ip地址,非此IP数据不接收',
  `tiger_shaped` char(32) NOT NULL DEFAULT '' COMMENT '握手符号',
  `heartbeat` int unsigned NOT NULL DEFAULT '0' COMMENT '上次心跳时间戳',
  `status` int NOT NULL DEFAULT '0' COMMENT '站点状态：0正常，非0为停用时间戳',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/
type Station struct {
	Id          int64          `json:"id" gorm:"primaryKey;autoIncrement;comment:主键id"`
	Name        string         `gorm:"type:varchar(20);column:name;not null;comment:站点名称" json:"name"`
	Location    datatypes.JSON `gorm:"type:json;column:location;not null;comment:站点所在经纬度" json:"location"`
	IpAddress   string         `gorm:"type:varchar(15);column:ip_address;not null;comment:站点ip地址,非此IP数据不接收" json:"ip_address"`
	TigerShaped string         `gorm:"type:char(32);column:tiger_shaped;not null;comment:握手符号" json:"tiger_shaped"`
	Heartbeat   int64          `gorm:"type:int(10);column:heartbeat;not null;comment:上次心跳时间戳" json:"heartbeat"`
	Status      int64          `gorm:"type:int(11);column:status;comment:站点状态：0正常，非0为停用时间戳" json:"status"`
	AddTime     time.Time      `gorm:"type:datetime;comment:创建时间" json:"add_time,omitempty"`
	UpdateTime  time.Time      `gorm:"type:datetime;comment:修改时间" json:"update_time,omitempty"`
}

// TableName 自定义表名
func (Station) TableName() string {
	return "pf_station_info"
}
