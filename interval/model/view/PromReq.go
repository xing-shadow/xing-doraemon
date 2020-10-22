/*
 * @Time : 2020/10/22 15:45
 * @Author : wangyl
 * @File : PromReq.go
 * @Software: GoLand
 */
package view

type GetProms struct {
	PaginationRequest
}

type CreateProm struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required"`
}

type ModifyProm struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type DeleteProm struct {
	ID uint `json:"id" binding:"required"`
}