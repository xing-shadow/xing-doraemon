package Resp

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func TestReflect(t *testing.T) {
	var resp Resp
	typ := reflect.TypeOf(&resp).Elem()
	val := reflect.ValueOf(&resp).Elem()
	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		fmt.Println(typeField, structField)
	}
}

func TestEqualFold(t *testing.T) {
	fmt.Println(strings.EqualFold("name", "Name"))
	fmt.Println(strings.EqualFold("name", "name"))
	fmt.Println(strings.EqualFold("name", "nAme"))
	fmt.Println(strings.EqualFold("name", "namE"))
	fmt.Println(strings.EqualFold("name", "1name"))
}

type Req struct {
	Page     uint `json:"page" query:"page1"`
	PageSize int  `json:"page_size" query:"size2"`
}

func TestBindParam(t *testing.T) {
	srv := gin.Default()
	srv.GET("/test", Handle(ParseQueryParam))
	srv.Run()
}

func ParseQueryParam(ctx *Context) {
	var req Req
	err := ctx.BindParam(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, req)
}

func TestA(t *testing.T) {
	m := md5.New()
	m.Write([]byte("123456"))
	fmt.Println(string(m.Sum(nil)))
}
