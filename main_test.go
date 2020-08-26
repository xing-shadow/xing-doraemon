/*
@Time : 2020/7/23 16:02
@Author : wangyl
@File : main_test.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	parseUrl, err := url.Parse("http://222.213.16.50:42001/dasdas/index.html")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(parseUrl.Path)
}
