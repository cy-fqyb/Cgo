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
        "/backend/common/onlineUserCount": {
            "get": {
                "description": "获取当前在线用户人数",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台公共接口"
                ],
                "summary": "获取当前在线用户人数",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.RT-int"
                        }
                    }
                }
            }
        },
        "/common/email": {
            "get": {
                "description": "发送邮箱验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端通用接口"
                ],
                "summary": "发送邮箱验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/common/test": {
            "get": {
                "description": "测试接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端通用接口"
                ],
                "summary": "测试接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/common/upload": {
            "post": {
                "description": "文件上传接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端通用接口"
                ],
                "summary": "文件上传接口",
                "parameters": [
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/front/msg/getMsgHistory": {
            "get": {
                "description": "获取用户消息历史记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端消息接口"
                ],
                "summary": "获取用户消息历史记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "好友id",
                        "name": "from_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.RT-models_MsgHistory"
                        }
                    }
                }
            }
        },
        "/front/msg/getMsgList": {
            "get": {
                "description": "获取用户消息列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端消息接口"
                ],
                "summary": "获取用户消息列表",
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
                            "$ref": "#/definitions/common.RT-models_MsgList"
                        }
                    }
                }
            }
        },
        "/front/msg/getMsgNum": {
            "get": {
                "description": "获取用户未读消息数量",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端消息接口"
                ],
                "summary": "获取用户未读消息数量",
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
                            "$ref": "#/definitions/common.RT-models_MsgNum"
                        }
                    }
                }
            }
        },
        "/front/room/applyJoinRoom": {
            "post": {
                "description": "申请加入房间",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端房间接口"
                ],
                "summary": "申请加入房间",
                "parameters": [
                    {
                        "description": "用户申请信息",
                        "name": "room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RoomApply"
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
        "/front/room/createRoom": {
            "post": {
                "description": "创建房间",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端房间接口"
                ],
                "summary": "创建房间",
                "parameters": [
                    {
                        "description": "房间信息",
                        "name": "room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Room"
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
        "/front/room/getRoomApply": {
            "get": {
                "description": "获取房间的申请信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端房间接口"
                ],
                "summary": "获取房间的申请信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "房主id",
                        "name": "master_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.RT-models_RoomApply"
                        }
                    }
                }
            }
        },
        "/front/room/getRoomUsers": {
            "get": {
                "description": "获取所有用户的信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端房间接口"
                ],
                "summary": "获取聊天群中的所有用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "群组id",
                        "name": "room_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.RT-models_Users"
                        }
                    }
                }
            }
        },
        "/front/room/handleRoomApplication": {
            "post": {
                "description": "处理用户加群申请",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端房间接口"
                ],
                "summary": "处理用户加群申请",
                "parameters": [
                    {
                        "description": "处理用户申请信息",
                        "name": "room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RoomApply"
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
        "/front/user/addFriend": {
            "post": {
                "description": "请求添加好友",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "请求添加好友",
                "parameters": [
                    {
                        "description": "处理申请的参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserFriend"
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
        "/front/user/deleteFriend": {
            "post": {
                "description": "删除好友",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "删除好友",
                "parameters": [
                    {
                        "description": "处理申请的参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserFriend"
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
                            "$ref": "#/definitions/models.UserRegister"
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
        "/front/user/searchUser": {
            "get": {
                "description": "搜索用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "搜索用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
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
        "/front/user/updateUserInfo": {
            "post": {
                "description": "修改用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前端用户接口"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "description": "用户信息",
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
        "common.RT-int": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "common.RT-models_MsgHistory": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/models.MsgHistory"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "common.RT-models_MsgList": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/models.MsgList"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "common.RT-models_MsgNum": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/models.MsgNum"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "common.RT-models_RoomApply": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/models.RoomApply"
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
        "models.MsgHistory": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "fromId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isRead": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "toId": {
                    "type": "string"
                }
            }
        },
        "models.MsgList": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "friendId": {
                    "type": "string"
                },
                "friendName": {
                    "type": "string"
                },
                "msgCount": {
                    "type": "integer"
                }
            }
        },
        "models.MsgNum": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "msgCount": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Room": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "master_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                }
            }
        },
        "models.RoomApply": {
            "type": "object",
            "required": [
                "master_id",
                "user_id"
            ],
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "master_id": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.UserFriend": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "friend_id": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.UserRegister": {
            "type": "object",
            "required": [
                "code",
                "email",
                "name",
                "password",
                "sex"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "code": {
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
	Host:             "159.75.155.236:8080",
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
