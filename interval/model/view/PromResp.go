/*
 * @Time : 2020/10/22 15:45
 * @Author : wangyl
 * @File : PromResp.go
 * @Software: GoLand
 */
package view

type PromList struct {
	PaginationResp
	PromList []PromItem `json:"prom_list"`
}

type PromItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
