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
	"strings"
	"time"
	"xing-doraemon/interval/app/HttpService/Handler"
	"xing-doraemon/interval/app/HttpService/middleware"
	"xing-doraemon/interval/service/CasbinService"
	"xing-doraemon/pkg/Utils"
	"xing-doraemon/pkg/app/Resp"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "xing-doraemon/docs"
	"xing-doraemon/global"
)

func Init() error {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.Cors())
	router.Use()
	// static file
	flag, err := Utils.IsFileExists("assets/build")
	if err != nil || !flag {
		panic("assets/dist not exist")
	}
	router.Static("/ant", "assets/build")
	router.Static("/static", "assets/build/static")
	router.StaticFile("/", "assets/build/index.html")
	router.NoRoute(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/api/") {
			ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		}
		ctx.File("assets/build")
	})
	LoginMiddleware := middleware.LoginAuth("/antd/login", middleware.RedirectTypeJson).Func()
	api := router.Group("/api/v1/")

	/*
		判断用户是否登录，并设置信息
	*/
	api.POST("/user/login", Resp.Handle(Handler.UserLogin))
	api.Use(LoginMiddleware)
	/*
		接口权限控制
	*/
	if global.GetAlterGatewayConfig().Casbin.Enable {
		api.Use(middleware.CasbinMiddleware(middleware.CasbinConfig{
			Enforcer: CasbinService.Casbin.SyncedEnforcer,
			Skipper: middleware.AllowPathPrefixSkipper([]string{
				"/api/v1/user/login",
			}...),
		}))
	}
	/*
		user
	*/
	{
		api.POST("/user/list", Resp.Handle(Handler.UserList))
		api.POST("/user/create", Resp.Handle(Handler.UserCreate))
		api.POST("/user/update", Resp.Handle(Handler.UserUpdate))
		api.POST("/user/delete", Resp.Handle(Handler.UserDelete))
	}
	/*
		rules
	*/
	{
		api.GET("/ruleId", Resp.Handle(Handler.GetRule))
		api.GET("/rule", Resp.Handle(Handler.GetRulePagination))
		api.POST("/rule", Resp.Handle(Handler.CreateRule))
		api.PUT("/rule", Resp.Handle(Handler.ModifyRule))
		api.DELETE("/rule", Resp.Handle(Handler.DeleteRule))
	}
	/*
		alerts
	*/
	{
		api.GET("/alerts", Resp.Handle(Handler.GetAlerts))
		api.POST("/alerts/confirm", Resp.Handle(Handler.ConfirmAlerts))

	}
	/*
		plans
	*/
	{
		api.GET("/planId", Resp.Handle(Handler.GetPlan))
		api.GET("/plan/allName", Resp.Handle(Handler.GetPlanAllName))
		api.GET("/plan", Resp.Handle(Handler.GetPlanPagination))
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
		api.GET("/prom/allName", Resp.Handle(Handler.GetPromAllName))
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
