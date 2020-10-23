/*
 * @Time : 2020/10/23 15:57
 * @Author : wangyl
 * @File : Alert.go
 * @Software: GoLand
 */
package AlertService

import (
	"encoding/json"
	"fmt"
	"xing-doraemon/interval/model/view"
)

func HandleAlert(input []byte) error {
	//TODO alert Handle
	var alerts view.Alerts
	err := json.Unmarshal(input, &alerts)
	if err != nil {
		return err
	} else {
		fmt.Println(alerts)
		return nil
	}
	/*todayZero, _ := time.ParseInLocation("2006-01-02", "2019-01-01 15:22:22", time.Local)
	for _, alert := range alerts {

	}*/
}
