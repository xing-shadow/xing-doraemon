/*
@Time : 2020/9/7 16:50
@Author : wangyl
@File : RuleHandler.go
@Software: GoLand
*/
package Handler

import (
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"xing-doraemon/gobal"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/model"
	"xing-doraemon/pkg/common"
)

type RuleHandler struct {
}

type Rule struct {
	Id    int64  `json:"id"`
	Expr  string `json:"expr"`
	Op    string `json:"op"`
	Value string `json:"value"`
	For   string `json:"for"`
	//Labels      map[string]string `json:"labels"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	PromId      int64  `json:"prom_id"`
	PlanId      int64  `json:"plan_id"`
}

func (r *RuleHandler) SendAllRules(ctx *gin.Context) {
	logger := gobal.GetLogger()
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logger.Errorf("Panic in SendAllRules:%v\n%s", e, buf)
		}
	}()
	var resp = &common.Res{}
	prom := ctx.Query("prom")
	id := ctx.Query("id")
	var receiver = &model.Rules{}
	results, err := receiver.Get(Invoker.GetDB(), prom, id)
	if err != nil {
		resp.Code = -1
		resp.Msg = ""
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	res := []Rule{}
	for _, i := range results {
		res = append(res, Rule{
			Id:          i.Id,
			Expr:        i.Expr,
			Op:          i.Op,
			Value:       i.Value,
			For:         i.For,
			Summary:     i.Summary,
			Description: i.Description,
			PromId:      i.Prom.Id,
			PlanId:      i.Plan.Id,
		})
	}
	resp.Code = 0
	resp.Msg = "Success"
	resp.Data = res
	ctx.JSON(http.StatusOK, resp)
}

func (*RuleHandler) AddRule(ctx *gin.Context) {
	logger := gobal.GetLogger()
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logger.Errorf("Panic in SendAllRules:%v\n%s", e, buf)
		}
	}()
	var ruleModel = &model.Rules{}
	var rule struct {
		Expr        string `json:"expr"`
		For         string `json:"for"`
		Op          string `json:"op"`
		Value       string `json:"value"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		PromId      int64  `json:"prom_id"`
		PlanId      int64  `json:"plan_id"`
	}
	var resp common.Res
	if err := ctx.BindJSON(&rule); err != nil {
		logger.Errorf("Unmarshal rule error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	ruleModel.Id = 0 //reset the "Id" to 0,which is very important:after a record is inserted,the value of "Id" will not be 0,but the auto primary key of the record
	ruleModel.Expr = rule.Expr
	ruleModel.Op = rule.Op
	ruleModel.Value = rule.Value
	ruleModel.For = rule.For
	ruleModel.Summary = rule.Summary
	ruleModel.Description = rule.Description
	ruleModel.Prom = &model.Proms{Id: rule.PromId} //cannot be models.RulesModel.Prom.Id=1,because the "Prom" is a pointer,which refers the null(cannot dereference the null pointer )
	ruleModel.Plan = &model.Plans{Id: rule.PlanId}
	err := ruleModel.InsertRule(Invoker.GetDB())
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, rule)
	ctx.JSON(http.StatusOK, resp)
}

func (*RuleHandler) UpdateRule(ctx *gin.Context) {
	logger := gobal.GetLogger()
	ruleId := ctx.Param("ruleid")
	var ruleModel = &model.Rules{}
	var rule struct {
		Expr        string `json:"expr"`
		Op          string `json:"op"`
		Value       string `json:"value"`
		For         string `json:"for"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		PromId      int64  `json:"prom_id"`
		PlanId      int64  `json:"plan_id"`
	}
	var resp common.Res
	if err := ctx.BindJSON(&rule); err != nil {
		logger.Errorf("Unmarshal rule error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	id, _ := strconv.ParseInt(ruleId, 10, 64)
	ruleModel.Id = id
	ruleModel.Expr = rule.Expr
	ruleModel.Op = rule.Op
	ruleModel.Value = rule.Value
	ruleModel.For = rule.For
	ruleModel.Description = rule.Description
	ruleModel.Summary = rule.Summary
	ruleModel.Prom = &model.Proms{Id: rule.PromId}
	ruleModel.Plan = &model.Plans{Id: rule.PlanId}
	err := ruleModel.UpdateRule(Invoker.GetDB())
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, ruleId)
	ctx.JSON(http.StatusOK, resp)
}

func (*RuleHandler) DeleteRule(ctx *gin.Context) {
	logger := gobal.GetLogger()
	ruleId := ctx.Param("ruleid")
	var Receiver = &model.Rules{}
	var resp common.Res
	err := Receiver.DeleteRule(Invoker.GetDB(), ruleId)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}
	logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, ruleId)
	ctx.JSON(http.StatusOK, resp)
}
