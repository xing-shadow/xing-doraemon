/*
@Time : 2020/7/20 9:52
@Author : wangyl
@File : main.go.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var HttpPort int

func main()  {
	for index, k := range os.Args {
		if k == "--http-port" {
			HttpPort,_ = strconv.Atoi(os.Args[index+1])
		}
	}
	path, _ := os.Executable()
	dir := filepath.Dir(path)
	router := gin.Default()
	fmt.Println(filepath.Join(dir,"/dist"))
	//router.Use(NetService.AllowFile())
	router.Use(static.Serve("/",static.LocalFile(filepath.Join(dir,"/dist"),true)))
	if err := router.Run(fmt.Sprintf(":%d",HttpPort)); err != nil {
		fmt.Println(err)
	}

}
