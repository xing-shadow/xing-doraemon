package HttpService

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"xing-doraemon/api/v1/Alert"
	"xing-doraemon/api/v1/Plan"
	"xing-doraemon/api/v1/Prom"
	"xing-doraemon/api/v1/Rule"
	"xing-doraemon/api/v1/User"
	"xing-doraemon/assets"
	"xing-doraemon/configs"
	_ "xing-doraemon/docs"
	"xing-doraemon/internal/app/HttpService/middleware"
	"xing-doraemon/pkg/App/Resp"
)

func Init(cfg configs.App) error {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// session init
	//router.Use(sessions.Sessions(UserService.SessionName, UserService.NewSessionStore()))
	//LoginMiddleware := middleware.LoginAuth("/antd/login").Func()

	router.Use(middleware.Cors())

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
	api.POST("/user/login", Resp.Handle(User.UserLogin))
	//user
	{
		api.GET("/user/list", Resp.Handle(User.UserList))
		api.POST("/user/create", Resp.Handle(User.UserCreate))
		api.POST("/user/update", Resp.Handle(User.UserUpdate))
		api.POST("/user/delete", Resp.Handle(User.UserDelete))
	}
	//rules
	{
		api.GET("/rule", Resp.Handle(Rule.GetRule))
		api.GET("/rules", Resp.Handle(Rule.GetRulePagination))
		api.POST("/rule/add", Resp.Handle(Rule.CreateRule))
		api.POST("/rule/update", Resp.Handle(Rule.ModifyRule))
		api.POST("/rule/delete", Resp.Handle(Rule.DeleteRule))
	}
	//alerts
	{
		api.GET("/alerts", Resp.Handle(Alert.GetAlerts))
		api.POST("/alerts/confirm", Resp.Handle(Alert.ConfirmAlerts))

	}
	//plans
	{
		api.GET("/plan", Resp.Handle(Plan.GetPlan))
		api.GET("/plan/allName", Resp.Handle(Plan.GetPlanAllName))
		api.GET("/plans", Resp.Handle(Plan.GetPlanPagination))
		api.POST("/plan/add", Resp.Handle(Plan.CreatePlan))
		api.POST("/plan/update", Resp.Handle(Plan.ModifyPlan))
		api.POST("/plan/delete", Resp.Handle(Plan.DeletePlan))

	}
	//proms
	{
		api.GET("/prom", Resp.Handle(Prom.GetProm))
		api.GET("/prom/allName", Resp.Handle(Prom.GetPromAllName))
		api.GET("/proms", Resp.Handle(Prom.GetPromsPagination))
		api.POST("/prom/add", Resp.Handle(Prom.CreateProm))
		api.POST("/prom/update", Resp.Handle(Prom.ModifyProm))
		api.POST("/prom/delete", Resp.Handle(Prom.DeleteProm))
	}
	service := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Httpport),
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	return service.ListenAndServe()
}
