/*
 * @Time : 2020/10/23 15:56
 * @Author : wangyl
 * @File : AlertReq.go
 * @Software: GoLand
 */
package view

type GetAlertsReq struct {
	PaginationRequest
}

type ConfirmAlertsReq struct {
	AlertList []uint `json:"alert_list"`
}
