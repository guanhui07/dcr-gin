package responseDto

import (
	"dcr-gin/app/model"
	"fmt"
	"time"
)

// ResponseUser 定义返回的数据模型
type ResponseUser struct {
	Id         int64      `json:"id"`
	UserName   string     `json:"username"`
	Status     int64      `json:"status"`
	AddTime    *time.Time `json:"add_time"`
	UpdateTime *time.Time `json:"update_time"`
}

// ToAccountModelToRes 将数据模型转换为返回值的
func ToAccountModelToRes(user model.User) ResponseUser {
	return ResponseUser{
		Id:       user.Id,
		UserName: user.UserName,
		Status:   user.Status,
	}
}

// ToAccountModelListToRes 列表的转换
func ToAccountModelListToRes(accountSlice []model.User) []ResponseUser {
	result := make([]ResponseUser, 0)
	for _, item := range accountSlice {
		fmt.Println(item.UserName)
		result = append(result, ToAccountModelToRes(item))
	}
	return result
}
