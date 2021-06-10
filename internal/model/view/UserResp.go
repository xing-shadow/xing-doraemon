/*
 * @Time : 2020/11/13 17:34
 * @Author : wangyl
 * @File : UserResp.go
 * @Software: GoLand
 */
package view

type UserListResp struct {
	PaginationResp
	UserList []UserItem `json:"user_list"`
}

type UserItem struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
}
