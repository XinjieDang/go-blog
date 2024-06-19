// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/public/login": {
            "post": {
                "description": "通过用户名和密码登录系统",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "认证"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录请求体",
                        "name": "loginUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功，返回用户信息及token",
                        "schema": {
                            "$ref": "#/definitions/response.LoginUserResp"
                        }
                    },
                    "400": {
                        "description": "请求参数无效",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "用户名或密码错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public/uploadFile": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "该接口用于获处理文件上传",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件上传"
                ],
                "summary": "通用文件上传-单文件",
                "responses": {
                    "200": {
                        "description": "成功返回文件信息，例如: 'xxxxx'",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/user/getHello": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "该接口用于获取当前登录用户的信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "测试接口",
                "responses": {
                    "200": {
                        "description": "成功返回用户信息，例如: 'John Doe'",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserInfoById/:id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "该接口用于获取当前登录用户的信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "通过id获取用户信息",
                "responses": {
                    "200": {
                        "description": "成功返回用户信息",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requests.LoginUserReq": {
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
        "response.LoginUserResp": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "last_login_time": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "response.SuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {}
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7001",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "blog API",
	Description:      "go 博客系统",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
