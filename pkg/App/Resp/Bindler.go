package Resp

import (
	"net/http"
)

type (
	// Binder is the interface that wraps the Bind method.
	Binder interface {
		BindParam(i interface{}) error
	}
)

// BindParam implements the `Binder#Bind` function.
func (ctx *Context) BindParam(i interface{}) error {
	switch ctx.Request.Method {
	case http.MethodGet:
		return ctx.BindQuery(i)
	default:
		if ctx.Binding == nil {
			return ctx.ShouldBindJSON(i)
		} else {
			return ctx.BindWith(i, ctx.Binding)
		}
	}
}
