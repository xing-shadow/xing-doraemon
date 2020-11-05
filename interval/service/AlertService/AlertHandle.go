/*
 * @Time : 2020/10/23 15:57
 * @Author : wangyl
 * @File : AlertHandle.go
 * @Software: GoLand
 */
package AlertService

import (
	"encoding/json"
	"fmt"
)

var (
	defaultHandle AlertHandle
)

type AlertHandle interface {
	HandleAlert(input []byte) error
}

func RegisterAlertHandle(alertHandle AlertHandle) {
	defaultHandle = alertHandle

}

func GetAlertHandle() AlertHandle {
	if defaultHandle == nil {
		return DefaultAlertHandle{}
	}
	return defaultHandle
}

type DefaultAlertHandle struct {
}

func (pThis DefaultAlertHandle) HandleAlert(input []byte) error {
	//TODO alert Handle
	var alerts []PromAlertItemList
	if err := json.Unmarshal(input, &alerts); err != nil {
		return err
	}
	fmt.Println(alerts)
	return nil
}
