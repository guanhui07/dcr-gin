package requestDto

type User struct {
	UserName string `json:"username" binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

/**
{
"username":"fsd",
"password":"sdf",
}
*/

type EditUser struct {
	UserName string `json:"username" binding:"min=2,max=10"`
	Status   int64  `json:"status"`
	Id       int32  `json:"id" binding:"required,gt=0"`
}

/**
{
"username":"fsd",
"status":"sdf",
"id":"sdf",
}
*/

type ChangeUser struct {
	UserName string `json:"username" binding:"min=2,max=10"`
	Status   int64  `json:"status"`
	Id       int32  `json:"id" binding:"required,gt=0"`
}

/**
{
"username":"fsd",
"status":"sdf",
"id":"sdf",
}
*/
type ChangeUserStatus struct {
	Status int64 `json:"status"`
	Id     int32 `json:"id" binding:"required,gt=0"`
}

/**
{
"status":"sdf",
"id":"sdf",
}
*/

type ChangePw struct {
	Id       int32  `json:"id" binding:"required,gt=0"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

/**
{
"password":"sdf",
"id":12,
}
*/

type ChangeAdminPw struct {
	//Id              int32  `json:"id" binding:"required,gt=0"`
	OldPassword    string `json:"old_password" binding:"required,min=6,max=16" `
	RepeatPassword string `json:"repeat_password" binding:"required,min=6,max=16"`
	Password       string `json:"password" binding:"required,min=6,max=16"`
}

/*
{
"old_password":"fsd",
"repeat_password":"sdf",
"password":"sdf",
}
*/

type UserList struct {
	Page     int `json:"page,default=1" form:"page,default=1"`
	PageRows int `json:"page_rows,default=10" form:"page_rows,default=10" `
}
