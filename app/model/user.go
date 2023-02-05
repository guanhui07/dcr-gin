package model

import "time"

// https://www.qetool.com/sql_json_go/sql.html   勾选 json gorm db
// https://www.qetool.com/sql_json_go/sql.html

/**
CREATE TABLE `pf_user_info` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(16) NOT NULL DEFAULT '' COMMENT '用户名',
  `pwd` char(64) NOT NULL DEFAULT '' COMMENT '密码的摘要',
  `salt` char(32) NOT NULL DEFAULT '' COMMENT '密码盐值',
  `last_ip` varchar(15) NOT NULL DEFAULT '0.0.0.0' COMMENT '最近一次登录IP',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '0为正常，非0为禁用时间戳',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息表';
*/
type User struct {
	Id         int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:主键id"`
	UserName   string    `gorm:"type:varchar(16);column:username;not null;unique;comment:账号" json:"username"`
	Password   string    `gorm:"type:char(32);column:pwd;not null;comment:账号密码" json:"pwd"`
	Salt       string    `gorm:"type:char(32);comment:密码盐值" json:"salt"`
	Status     int64     `gorm:"type:int(11);comment:0为正常，非0为禁用时间戳" json:"status"`
	LastIp     string    `gorm:"type:varchar(16);column:last_ip;" json:"last_ip"`
	AddTime    time.Time `gorm:"type:datetime;comment:创建时间" json:"add_time,omitempty"`
	UpdateTime time.Time `gorm:"type:datetime;comment:修改时间" json:"update_time,omitempty"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "pf_user_info"
}
