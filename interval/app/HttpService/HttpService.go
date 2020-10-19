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
	"github.com/gin-gonic/gin"

	"xing-doraemon/gobal"
)

func Init() error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowedOrigins:   []string{"http://localhost", "http://127.0.0.1"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}))
	//store := sessions.NewCookieStore([]byte("xing-shadow12345"))
	//store.Options(sessions.Options{
	//	Path:     "./tmp",
	//	MaxAge:   3600,
	//	Secure:   false,
	//	HttpOnly: true,
	//})
	//router.Use(sessions.Sessions("gosessionid", store))
	api := router.Group("/api/v1/")
	/*
		rules
	*/
	{
		api.GET("/rules")
		api.POST("/rule")
		api.DELETE("/rule/:ruleid")
	}
	/*
		alerts
	*/
	{
		api.GET("/alerts")
		api.GET("/alerts/rules/:ruleid")
		api.GET("/alerts/classify")
		api.PUT("/alerts")
		api.POST("/alerts")
	}
	/*
		plans
	*/
	{
		api.GET("/plans")
		api.POST("/plans")
		api.GET("/plans/:planid/receivers")
		api.POST("/plans/:planid/receivers")
		api.PUT("/plans/:planid")
		api.DELETE("/plans/:planid")

	}
	/*
		proms
	*/
	{
		api.GET("/proms")
		api.POST("/proms")
		api.PUT("/proms/:id")
		api.DELETE("/proms/:id")
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
