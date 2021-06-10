/*
@Time : 2020/9/3 16:46
@Author : wangyl
@File : HttpService.go
@Software: GoLand
*/
package HttpService

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/fs"
	"net/http"
	"strings"
	"time"
	"xing-doraemon/api/v1/Alert"
	"xing-doraemon/api/v1/Plan"
	"xing-doraemon/api/v1/Prom"
	"xing-doraemon/api/v1/Rule"
	"xing-doraemon/api/v1/User"
	"xing-doraemon/assets"
	"xing-doraemon/configs"
	_ "xing-doraemon/docs"
	"xing-doraemon/internal/app/HttpService/middleware"
	UserService "xing-doraemon/internal/service/UserService"
	"xing-doraemon/pkg/App/Resp"
)

func Init(cfg configs.App) error {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// session init
	router.Use(sessions.Sessions(UserService.SessionName, UserService.NewSessionStore()))

	router.Use(middleware.Cors())
	//LoginMiddleware := middleware.LoginAuth("/antd/login").Func()

	// static file
	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, http.DetectContentType(assets.GetIndex()), assets.GetIndex())
	})
	assetFS, err := fs.Sub(assets.GetStaticAssets(), "build/static")
	if err != nil {
		return err
	}
	router.StaticFS("/static", http.FS(assetFS))
	router.NoRoute(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/api/") {
			ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		data, err := assets.GetStaticAssets().ReadFile("build" + ctx.Request.URL.Path)
		if err != nil {
			ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		ctx.Data(http.StatusOK, http.DetectContentType(data), data)
	})

	api := router.Group("/api/v1/")
	/*
		判断用户是否登录，并设置信息
	*/
	api.POST("/user/login", Resp.Handle(User.UserLogin))
	//api.Use(LoginMiddleware)

	//user
	{
		api.POST("/user/list", Resp.Handle(User.UserList))
		api.POST("/user/create", Resp.Handle(User.UserCreate))
		api.POST("/user/update", Resp.Handle(User.UserUpdate))
		api.POST("/user/delete", Resp.Handle(User.UserDelete))
	}
	//rules
	{
		api.GET("/ruleId", Resp.Handle(Rule.GetRule))
		api.GET("/rule", Resp.Handle(Rule.GetRulePagination))
		api.POST("/rule", Resp.Handle(Rule.CreateRule))
		api.PUT("/rule", Resp.Handle(Rule.ModifyRule))
		api.DELETE("/rule", Resp.Handle(Rule.DeleteRule))
	}
	//alerts
	{
		api.GET("/alerts", Resp.Handle(Alert.GetAlerts))
		api.POST("/alerts/confirm", Resp.Handle(Alert.ConfirmAlerts))

	}
	//plans
	{
		api.GET("/planId", Resp.Handle(Plan.GetPlan))
		api.GET("/plan/allName", Resp.Handle(Plan.GetPlanAllName))
		api.GET("/plan", Resp.Handle(Plan.GetPlanPagination))
		api.POST("/plan", Resp.Handle(Plan.CreatePlan))
		//TODO
		//api.GET("/plan/:planId/rules")
		api.PUT("/plan", Resp.Handle(Plan.ModifyPlan))
		api.DELETE("/plan", Resp.Handle(Plan.DeletePlan))

	}
	//proms
	{
		api.GET("/promId", Resp.Handle(Prom.GetProm))
		api.GET("/prom/allName", Resp.Handle(Prom.GetPromAllName))
		api.GET("/prom", Resp.Handle(Prom.GetPromsPagination))
		api.POST("/prom", Resp.Handle(Prom.CreateProm))
		api.PUT("/prom", Resp.Handle(Prom.ModifyProm))
		api.DELETE("/prom", Resp.Handle(Prom.DeleteProm))
	}
	service := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Httpport),
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	return service.ListenAndServe()
}
