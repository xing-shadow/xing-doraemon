/*
 * @Time : 2020/11/17 17:33
 * @Author : wangyl
 * @File : CasbinService_test.go
 * @Software: GoLand
 */
package CasbinService

import (
	"github.com/casbin/casbin/v2"
	"testing"
)

func TestCasbinPermission(t *testing.T) {
	model := "D:\\go\\src\\xing-doraemon\\internal\\service\\CasbinService\\model.conf"
	e, err := casbin.NewSyncedEnforcer()
	if err != nil {
		t.Fatal(err)
	}
	policy := "D:\\go\\src\\xing-doraemon\\internal\\service\\CasbinService\\policy.csv"
	err = e.InitWithFile(model, policy)
	if err != nil {
		t.Fatal(err)
	}
	e.EnableEnforce(true)
	e.EnableLog(true)
	testData := [][]string{
		/*
			expect:true
		*/
		{
			"xing", "/hello", "get",
		},
		{
			"xing", "/hello", "delete",
		},
		{
			"tom", "/hello", "post",
		},
		/*
			expect:false
		*/
		{
			"tom", "/hello", "get",
		},
		{
			"tom", "/hello", "delete",
		},
		{
			"xing", "/hello", "post",
		},
	}
	for _, datum := range testData {
		_, err := e.Enforce(datum[0], datum[1], datum[2])
		if err != nil {
			t.Fatal(err)
		}
	}
}
