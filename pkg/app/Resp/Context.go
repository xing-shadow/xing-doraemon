package Resp

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
}

type HandFunc func(ctx *Context)

func Handle(handFunc HandFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctxNew := &Context{ctx}
		handFunc(ctxNew)
	}
}

type Response struct {
	Status int `json:"-"`

	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type RespOption func(resp *Response)

func (ctx *Context) ToResponse(code int, msg string, options ...RespOption) {
	result := new(Response)
	result.Code = code
	result.Message = msg
	for _, opt := range options {
		opt(result)
	}
	ctx.JSON(result.Status, result)
}

func (ctx *Context) WithStatus(status int) RespOption {
	return func(resp *Response) {
		resp.Status = status
	}
}

func (ctx *Context) WithData(data interface{}) RespOption {
	return func(resp *Response) {
		resp.Data = data
	}
}

func (ctx *Context) Bind(i interface{})  {

}