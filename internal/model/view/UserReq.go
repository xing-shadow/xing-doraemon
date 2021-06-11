package view

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserListReq struct {
	Page     uint `json:"page"`
	PageSize uint `json:"page_size"`
}

type UserCreateReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserUpdateReq struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserDeleteReq struct {
	Id uint `json:"id"`
}
