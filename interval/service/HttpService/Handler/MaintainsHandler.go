/*
@Time : 2020/9/7 19:46
@Author : wangyl
@File : MaintainsHandler.go
@Software: GoLand
*/
package Handler

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"xing-doraemon/gobal"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/model"
	"xing-doraemon/pkg/common"
)

type MaintainHandler struct {
}

func (*MaintainHandler) GetAllMaintains(ctx *gin.Context) {
	var Maintain = &model.Maintains{}
	maintains := Maintain.GetAllMaintains(Invoker.GetDB())
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: maintains,
	})
}

func (*MaintainHandler) GetHosts(ctx *gin.Context) {
	var Host = &model.Hosts{}
	mid := ctx.Param(":id")
	hosts := Host.GetHosts(Invoker.GetDB(), mid)
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: hosts,
	})
}

func (*MaintainHandler) AddMaintain(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var data struct {
		Flag      bool   `json:"flag"`
		TimeStart string `json:"time_start"`
		TimeEnd   string `json:"time_end"`
		Month     string `json:"month"`
		DayStart  int8   `json:"day_start"`
		DayEnd    int8   `json:"day_end"`
		Valid     string `json:"valid"`
		Hosts     string `json:"hosts"`
	}
	var resp common.Res
	err := ctx.BindJSON(&data)
	if err != nil {
		logger.Errorf("Unmarshal prom error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		var maintain model.Maintains
		maintain.TimeStart = data.TimeStart
		maintain.TimeEnd = data.TimeEnd
		maintain.DayStart = data.DayStart
		maintain.DayEnd = data.DayEnd
		if data.TimeStart > data.TimeEnd {
			maintain.Flag = true
		} else {
			maintain.Flag = false
		}
		validTime, _ := time.ParseInLocation("2006-01-02 15:04:05", data.Valid, time.Local)
		maintain.Valid = &validTime
		monthList := strings.Split(data.Month, ",")
		for _, m := range monthList {
			e, _ := strconv.ParseFloat(m, 64)
			maintain.Month = maintain.Month | int(math.Pow(2, e))
		}
		err = maintain.AddMaintains(Invoker.GetDB(), data.Hosts)
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Info("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, data)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*MaintainHandler) UpdateMaintain(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var data struct {
		Flag      bool   `json:"flag"`
		TimeStart string `json:"time_start"`
		TimeEnd   string `json:"time_end"`
		Month     string `json:"month"`
		DayStart  int8   `json:"day_start"`
		DayEnd    int8   `json:"day_end"`
		Valid     string `json:"valid"`
		Hosts     string `json:"hosts"`
	}
	var resp common.Res
	mid := ctx.Param("id")
	id, _ := strconv.ParseInt(mid, 10, 64)
	err := ctx.BindJSON(&data)
	if err == nil {
		var maintain model.Maintains
		maintain.Id = id
		maintain.TimeStart = data.TimeStart
		maintain.TimeEnd = data.TimeEnd
		maintain.DayStart = data.DayStart
		maintain.DayEnd = data.DayEnd
		if data.TimeStart > data.TimeEnd {
			maintain.Flag = true
		} else {
			maintain.Flag = false
		}
		validTime, _ := time.ParseInLocation("2006-01-02 15:04:05", data.Valid, time.Local)
		maintain.Valid = &validTime
		monthList := strings.Split(data.Month, ",")
		for _, m := range monthList {
			e, _ := strconv.ParseFloat(m, 64)
			maintain.Month = maintain.Month | int(math.Pow(2, e))
		}
		//fmt.Println([]byte(data.Hosts))
		err = maintain.UpdateMaintains(Invoker.GetDB(), data.Hosts)
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Info("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, data)
	} else {
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*MaintainHandler) DeleteMaintain(ctx *gin.Context) {
	logger := gobal.GetLogger()
	id := ctx.Param(":id")
	var maintain = &model.Maintains{}
	var resp common.Res
	err := maintain.DeleteMaintains(Invoker.GetDB(), id)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Info("%s %s %s %s", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, id)
	ctx.JSON(http.StatusOK, resp)
}
