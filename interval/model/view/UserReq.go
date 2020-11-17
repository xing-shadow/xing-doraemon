/*
 * @Time : 2020/11/13 17:33
 * @Author : wangyl
 * @File : UserReq.go
 * @Software: GoLand
 */
package view

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserListReq struct {
	PaginationRequest
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
