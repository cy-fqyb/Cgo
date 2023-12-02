// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/front/common/ws/{id}": {
            "get": {
                "description": "websocket对话通信接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "websocket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "消息体",
                        "name": "msg",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.WsDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/front/user/getUserFriendApply": {
            "get": {
                "description": "获取用户好友申请",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "获取用户好友申请",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Result"
                        }
                    }
                }
            }
        },
        "/front/user/getUserFriends": {
            "get": {
                "description": "获取用户好友",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "获取用户好友",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.RT-array_models_Users"
                        }
                    }
                }
            }
        },
        "/front/user/handleApplication": {
            "post": {
                "description": "处理用户好友申请",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "处理用户好友申请",
                "parameters": [
                    {
                        "description": "处理申请的参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserApplyDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.RT-string"
                        }
                    }
                }
            }
        },
        "/front/user/login": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "Id",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.RT-models_Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/front/user/register": {
            "post": {
                "description": "用户注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "description": "注册用户的信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.RT-string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.RT-array_models_Users": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Users"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "common.RT-models_Users": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/models.Users"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "common.RT-string": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "common.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "errs": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "dto.UserApplyDto": {
            "type": "object",
            "required": [
                "is_accept",
                "user_id"
            ],
            "properties": {
                "friend_id": {
                    "type": "string"
                },
                "is_accept": {
                    "type": "boolean"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.UserLoginDto": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.WsDto": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "to_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.Users": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "sex"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "synopsis": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0 版本",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Im在线通信系统接口文档",
	Description:      "im在线通信系统前后台接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
