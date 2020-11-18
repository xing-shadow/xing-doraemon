// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/alerts": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取告警列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页号",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/alerts/confirm": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "告警确认",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ConfirmAlertsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/plan": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取Plan列表，分页",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页序号",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页大小",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改plan",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ModifyPlanReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "创建plan",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.CreatePlanReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除plan",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.DeleteRuleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/plan/allNames": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取所有prom名",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/planID": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取单个Plan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "序号",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/prom": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取proms, 分页",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页序号",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页大小",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改prom",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ModifyProm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "创建prom",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.CreateProm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除prom",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.DeleteProm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/prom/allName": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取所有prom名",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/promId": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取单个prom",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页序号",
                        "name": "Id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/rule": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取rules列表，分页",
                "parameters": [
                    {
                        "type": "string",
                        "description": "序号",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "序号",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改rule",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ModifyRuleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除rule",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.DeleteRuleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/ruleId": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取单个rule",
                "parameters": [
                    {
                        "type": "string",
                        "description": "序号",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/create": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.UserCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/delete": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.UserDeleteReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "序号",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "序号",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/update": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.UserUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Resp.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Resp.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "view.ConfirmAlertsReq": {
            "type": "object",
            "properties": {
                "alert_list": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "view.CreatePlanReq": {
            "type": "object",
            "required": [
                "period"
            ],
            "properties": {
                "end_time": {
                    "type": "string"
                },
                "expression": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "period": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "view.CreateProm": {
            "type": "object",
            "required": [
                "name",
                "url"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "view.CreateRuleReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "expr": {
                    "type": "string"
                },
                "for": {
                    "description": "持续时间，单位秒",
                    "type": "integer"
                },
                "op": {
                    "type": "string"
                },
                "plan_name": {
                    "type": "string"
                },
                "prom_name": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "view.DeleteProm": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "view.DeleteRuleReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "view.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "view.ModifyPlanReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "end_time": {
                    "type": "string"
                },
                "expression": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "period": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "view.ModifyProm": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "view.ModifyRuleReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "expr": {
                    "type": "string"
                },
                "for": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "op": {
                    "type": "string"
                },
                "plan_name": {
                    "type": "string"
                },
                "prom_name": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "view.UserCreateReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "view.UserDeleteReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "view.UserUpdateReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "2.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "prometheus Alert management center",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
