/*
@Time : 2020/9/3 16:46
@Author : wangyl
@File : HttpService.go
@Software: GoLand
*/
package HttpService

import (
	"fmt"
	"net/http"
	"time"
	"xing-doraemon/interval/app/HttpService/Handler"
	"xing-doraemon/pkg/App/Resp"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "xing-doraemon/docs"
	"xing-doraemon/global"
)

func Init() error {
	router := gin.Default()
	//store := sessions.NewCookieStore([]byte("xing-shadow12345"))
	//store.Options(sessions.Options{
	//	Path:     "./tmp",
	//	MaxAge:   3600,
	//	Secure:   false,
	//	HttpOnly: true,
	//})
	//router.Use(sessions.Sessions("gosessionid", store))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(Cors())
	router.Use(func(ctx *gin.Context) {
		ctx.Next()
		for _, h := range ctx.Writer.Header() {
			fmt.Println(h)
		}
	})
	api := router.Group("/api/v1/")
	/*
		rules
	*/
	{
		api.GET("/ruleID", Resp.Handle(Handler.GetRule))
		api.GET("/rule", Resp.Handle(Handler.GetRulePagination))
		api.GET("/rules", Resp.Handle(Handler.GetAllRule))
		api.POST("/rule", Resp.Handle(Handler.CreateRule))
		api.PUT("/rule", Resp.Handle(Handler.ModifyRule))
		api.DELETE("/rule", Resp.Handle(Handler.DeleteRule))
	}
	/*
		alerts
	*/
	{
		/*	api.GET("/alerts")
			api.GET("/alerts/rules/:ruleId")
			api.GET("/alerts/classify")
			api.PUT("/alerts")
			api.POST("/alerts")*/
	}
	/*
		plans
	*/
	{
		api.GET("/planID", Resp.Handle(Handler.GetPlan))
		api.GET("/plan", Resp.Handle(Handler.GetPlanPagination))
		api.GET("/plans", Resp.Handle(Handler.GetAllPlan))
		api.POST("/plan", Resp.Handle(Handler.CreatePlan))
		//TODO
		//api.GET("/plan/:planId/rules")
		api.PUT("/plan", Resp.Handle(Handler.ModifyPlan))
		api.DELETE("/plan", Resp.Handle(Handler.DeletePlan))

	}
	/*
		proms
	*/
	{
		api.GET("/promId", Resp.Handle(Handler.GetProm))
		api.GET("/prom", Resp.Handle(Handler.GetPromsPagination))
		api.POST("/prom", Resp.Handle(Handler.CreateProm))
		api.PUT("/prom", Resp.Handle(Handler.ModifyProm))
		api.DELETE("/prom", Resp.Handle(Handler.DeleteProm))
	}
	service := http.Server{
		Addr:         fmt.Sprintf(":%d", global.GetAlterGatewayConfig().App.Httpport),
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	go func() {
		if err := service.ListenAndServe(); err != nil {
			global.GetLogger().Panic("start Service fail: ", err)
		}
	}()
	return nil
}
