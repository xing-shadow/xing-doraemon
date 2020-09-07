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

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"xing-doraemon/gobal"
	"xing-doraemon/interval/service/HttpService/Handler"
	"xing-doraemon/pkg/middleware/HttpMiddleware"
)

func InitHttpService() error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowedOrigins:   []string{"http://10.*.*.*:*", "http://localhost:*", "http://127.0.0.1:*", "http://172.*.*.*:*", "http://192.*.*.*:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*", "content-time"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}))
	store := sessions.NewCookieStore([]byte("xing-shadow12345"))
	store.Options(sessions.Options{
		Path:     "./tmp",
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("gosessionid", store))
	api := router.Group("/api/v1/", HttpMiddleware.FilterUserMiddleware())
	/*
		login
	*/
	{
		loginHandler := &Handler.LoginHandler{}
		api.GET("/login/method", loginHandler.GetMethod)
		api.GET("/login/username", loginHandler.Username)
		api.POST("/login/local", loginHandler.Local)
		api.POST("/login/ldap", loginHandler.Ldap)
	}
	/*
		logout
	*/
	{
		logoutHandler := &Handler.LogoutHandler{}
		api.GET("/logout", logoutHandler.Logout)
	}
	/*
		user
	*/
	{
		userHandler := &Handler.UserHandler{}
		api.GET("/users", userHandler.GetAllUsers)
		api.POST("/users", userHandler.AddUser)
		api.PUT("/users", userHandler.UpdatePassword)
		api.DELETE("/users/:id", userHandler.DeleteUsers)
	}
	/*
		rules
	*/
	{
		ruleHandler := &Handler.RuleHandler{}
		api.GET("/rules", ruleHandler.SendAllRules)
		api.POST("/rules", ruleHandler.AddRule)
		api.PUT("/rules/:ruleid", ruleHandler.UpdateRule)
		api.DELETE("/rules/:ruleid", ruleHandler.DeleteRule)
	}
	/*
		alerts
	*/
	{ // TODO write code
		alertHandler := &Handler.AlertHandler{}
		api.GET("/alerts", alertHandler.GetAlerts)
		api.GET("/alerts/:ruleid", alertHandler.ShowAlerts)
		api.GET("/alerts/classify", alertHandler.ClassifyAlerts)
		api.PUT("/alerts", alertHandler.Confirm)
		api.POST("/alerts", alertHandler.HandleAlerts)
	}
	/*
		plans
	*/
	{
		planHandler := &Handler.PlanHandler{}
		api.GET("/plans", planHandler.GetAllPlans)
		api.POST("/plans", planHandler.AddPlan)
		api.GET("/plans/:planid/receivers", planHandler.GetAllReceiver)
		api.GET("/plans/:planid/receivers", planHandler.AddReceiver)
		api.PUT("/plans/:planid", planHandler.UpdatePlan)
		api.GET("/plans/:planid", planHandler.DeletePlan)

	}
	/*
		receivers
	*/
	{
		receiverHandler := &Handler.ReceiversHandler{}
		api.PUT("/receivers/:receiverid", receiverHandler.UpdateReceiver)
		api.DELETE("/receivers/:receiverid", receiverHandler.DeleteReceiver)
	}
	/*
		groups
	*/
	{
		groupHandler := &Handler.GroupHandler{}
		api.GET("/groups", groupHandler.GetAllGroup)
		api.POST("/groups", groupHandler.AddGroup)
		api.PUT("/groups/:id", groupHandler.UpdateGroup)
		api.DELETE("/groups/:id", groupHandler.DeleteGroup)
	}
	/*
		proms
	*/
	{
		promsHandler := &Handler.PromsHandler{}
		api.GET("/proms", promsHandler.GetAllProms)
		api.POST("/proms", promsHandler.AddProm)
		api.PUT("/proms/:id", promsHandler.UpdateProm)
		api.DELETE("/proms/:id", promsHandler.DeleteProm)
	}
	/*
		maintains
	*/
	{
		maintainHandler := &Handler.MaintainHandler{}
		api.GET("/maintains", maintainHandler.GetAllMaintains)
		api.GET("/maintains/:id/hosts", maintainHandler.GetHosts)
		api.POST("/maintains", maintainHandler.AddMaintain)
		api.PUT("/maintains/:id", maintainHandler.UpdateMaintain)
		api.DELETE("/maintains/:id", maintainHandler.DeleteMaintain)
	}
	/*
		manages
	*/
	{
		managerHandler := &Handler.ManagerHandler{}
		api.GET("/manages", managerHandler.GetAll)
		api.POST("/manages", managerHandler.AddManage)
		api.PUT("/manages/:id", managerHandler.UpdateManage)
		api.DELETE("/manages/:id", managerHandler.DeleteManage)
	}
	/*
		configs

	*/
	{
		configHandler := &Handler.ConfigHandler{}
		api.GET("/configs", configHandler.GetAll)
		api.POST("/configs", configHandler.AddConfig)
		api.PUT("/configs/:id", configHandler.UpdateConfig)
		api.DELETE("/configs/:id", configHandler.DeleteConfig)
	}
	service := http.Server{
		Addr:         fmt.Sprintf(":%d", gobal.GetAlterGatewayConfig().App.Httpport),
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	go func() {
		if err := service.ListenAndServe(); err != nil {
			gobal.GetLogger().Panic("start Service fail: ", err)
		}
	}()
	return nil
}
