package service

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/requestDto"
	"dcr-gin/app/responseDto"
	"dcr-gin/app/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type UserService struct{}

func findUser(c *gin.Context, userParams requestDto.User) (userResp responseDto.ResponseUser, err error) {
	result := global.DB.Model(&model.User{}).First(&userParams)
	if result.RowsAffected > 0 {
		// 存在记录
		return userResp, nil
	}
	return userResp, errors.New("用户不存在！")

}

// AddUser 新增用户
func AddUser(c *gin.Context, userParams requestDto.User) error {
	userModel := model.User{}
	if !errors.Is(global.DB.Where("username = ?", userParams.UserName).
		First(&userModel).Error, gorm.ErrRecordNotFound) {
		// 判断用户名是否存在
		return errors.New("账号已存在")
	}
	newPassword, salt, err := utils.GeneratePassword(userParams.Password)
	if err != nil {
		return errors.New("密码加密错误")
	}

	fmt.Println(newPassword)
	newUser := model.User{
		UserName: userParams.UserName,
		Password: newPassword,
		Salt:     salt,
		Status:   1,
		AddTime:  time.Now(),
	}
	// 插入 新增保存到db
	ret := global.DB.Create(&newUser)
	return ret.Error
}

// EditUser 编辑用户
func EditUser(c *gin.Context, userParams requestDto.EditUser) error {
	userModel := model.User{}
	//判断db是否存在
	db := global.DB
	if errors.Is(db.Where("id=?", userParams.Id).First(&userModel).Error, gorm.ErrRecordNotFound) {
		//记录不存在
		return errors.New("此账号不存在")
	}

	if userParams.Status != 0 {
		//获取unix时间戳
		userParams.Status = time.Now().Unix()
	}
	userModel.Status = userParams.Status
	userModel.UserName = userParams.UserName
	userModel.UpdateTime = time.Now()
	// 更新db
	ret := db.Where("id=?", userParams.Id).
		Select("status", "username", "update_time").
		Updates(userModel)
	return ret.Error
}

// ChangeStatus 修改状态
func ChangeStatus(c *gin.Context, UserParams requestDto.ChangeUserStatus) error {
	userModel := model.User{}
	if errors.Is(global.DB.Where("id=?", UserParams.Id).
		First(&userModel).Error, gorm.ErrRecordNotFound) {
		//记录不存在 查找一行
		return errors.New("此账号不存在")
	}
	if UserParams.Status != 0 {
		UserParams.Status = time.Now().Unix()
	}
	userModel.Status = UserParams.Status
	userModel.UpdateTime = time.Now()
	// 更新db
	ret := global.DB.Model(&userModel).
		Select("status", "update_time").
		Updates(userModel)
	//info := global.DB.Where("id=?",UserParams.Id).Select("status","update_time").Updates(userModel)
	return ret.Error
}

// ChangePw 修改密码
func ChangePw(c *gin.Context, userParams requestDto.ChangePw) error {
	userModel := model.User{}
	if errors.Is(global.DB.Where("id=?", userParams.Id).First(&userModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("此账号不存在")
	}
	newPw, salt, err := utils.GeneratePassword(userParams.Password)
	if err != nil {
		return errors.New("密码加密错误")
	}
	userModel.Password = newPw
	userModel.Salt = salt
	userModel.UpdateTime = time.Now()
	dump.P(newPw)
	// 更新db
	ret := global.DB.Where("id=?", userParams.Id).
		Select("pwd", "salt", "update_time").
		Save(userModel)
	return ret.Error
}

// ChangeAdminPw 修改管理员密码
func ChangeAdminPw(c *gin.Context, userParams requestDto.ChangeAdminPw) error {
	userModel := model.User{}
	err := global.DB.Where("username=?", "admin").First(&userModel).Error

	if err != nil {
		return errors.New("管理员账号获取失败")
	}
	result, _ := utils.CheckPassword(userModel.Password, userParams.OldPassword, userModel.Salt)
	if result == false {
		return errors.New("密码错误，无法修改密码")
	}
	if userParams.Password != userParams.RepeatPassword {
		return errors.New("两次新密码输入不一致")
	}
	newPw, salt, err := utils.GeneratePassword(userParams.Password)
	if err != nil {
		return errors.New("密码加密错误")
	}
	userModel.Password = newPw
	userModel.Salt = salt
	userModel.UpdateTime = time.Now()
	// 更新db
	ret := global.DB.Where("username=?", "admin").
		Select("pwd", "salt", "update_time").
		Updates(userModel)
	return ret.Error
}

type ResponseUsers struct {
	Id       int64     `json:"id"`
	UserName string    `json:"username"`
	LastIp   string    `json:"last_ip"`
	Status   int64     `json:"status"`
	AddTime  time.Time `json:"add_time"`
}

//type UserData struct {
//	ResponseUser []Response_users `json:"user"`
//}

type Response struct {
	CurrentPage int             `json:"current_page"`
	PageRows    int             `json:"page_rows"`
	TotalCount  int64           `json:"total_count"`
	Data        []ResponseUsers `json:"data"`
}

// UserList 用户列表
func UserList(c *gin.Context, userListData requestDto.UserList) (err error, ConfigInfo Response) {
	var userListCount []model.User
	userModelList := []model.User{}

	var responseData Response
	page := userListData.Page
	pageRows := userListData.PageRows

	var count int64
	// 查询总条数
	global.DB.Find(&userListCount).Count(&count)
	// 查询列表
	if err := global.DB.Limit(pageRows).
		Offset((page - 1) * pageRows).
		Find(&userModelList).Error; err != nil {
		return errors.New("获取用户列表失败"), responseData
	}
	var Data []ResponseUsers
	copier.Copy(&Data, userModelList)

	responseData = Response{
		CurrentPage: page,
		PageRows:    pageRows,
		TotalCount:  count,
		Data:        Data,
	}
	return err, responseData

}

// ChangePassword @function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error
func ChangePassword(id uint, newPassword string) (userResp *responseDto.ResponseUser, err error) {
	var userModel model.User
	if err = global.DB.Where("id = ?", id).First(&userModel).Error; err != nil {
		//记录不存在
		return nil, err
	}
	newPassword, salt, err := utils.GeneratePassword(newPassword)
	if err != nil {
		return nil, errors.New("密码加密错误")
	}
	userModel.Password = newPassword
	userModel.Salt = salt
	err = global.DB.Save(&userModel).Error
	return nil, err
}
