package view

type UserListResp struct {
	Pagination `json:"pagination"`
	UserList   []UserItem `json:"list"`
}

type UserItem struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
}
