package view

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserListReq struct {
	Page     uint `query:"page"`
	PageSize uint `query:"page_size"`
}

type UserCreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdateReq struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDeleteReq struct {
	Id uint `json:"id"`
}
