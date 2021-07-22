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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/repo/file": {
            "get": {
                "description": "获取仓库具体路径下的内容",
                "tags": [
                    "仓库"
                ],
                "summary": "获取仓库具体路径下的内容",
                "parameters": [
                    {
                        "type": "string",
                        "description": "access_token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "login",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文件UID",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功后返回值",
                        "schema": {
                            "$ref": "#/definitions/response.FileInfo"
                        }
                    }
                }
            },
            "put": {
                "description": "更新文件",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "仓库"
                ],
                "summary": "更新文件",
                "parameters": [
                    {
                        "description": "FileInfo",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.FileInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功后返回值",
                        "schema": {
                            "$ref": "#/definitions/response.FileInfo"
                        }
                    }
                }
            },
            "post": {
                "description": "新建文件",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "仓库"
                ],
                "summary": "新建文件",
                "parameters": [
                    {
                        "description": "FileInfo",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.FileInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功后返回值",
                        "schema": {
                            "$ref": "#/definitions/response.FileInfo"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除文件",
                "tags": [
                    "仓库"
                ],
                "summary": "删除文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "access_token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "login",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文件UID",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文件sha",
                        "name": "sha",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功后返回值",
                        "schema": {
                            "$ref": "#/definitions/response.FileInfo"
                        }
                    }
                }
            }
        },
        "/repo/info": {
            "get": {
                "description": "获取仓库信息，没有则初始化仓库，并返回.shownote/workspace.json的内容",
                "tags": [
                    "仓库"
                ],
                "summary": "获取仓库信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "access_token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "login",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.FileInfo"
                        }
                    }
                }
            }
        },
        "/repo/upload": {
            "post": {
                "description": "上传文件",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "仓库"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件类型,image|viodo|audio",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功后返回值",
                        "schema": {
                            "$ref": "#/definitions/response.FileInfo"
                        }
                    }
                }
            }
        },
        "/repo/workspace": {
            "get": {
                "description": "获取workspace内容",
                "tags": [
                    "仓库"
                ],
                "summary": "获取workspace内容",
                "parameters": [
                    {
                        "type": "string",
                        "description": "access_token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "login",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功后返回值",
                        "schema": {
                            "$ref": "#/definitions/response.FileInfo"
                        }
                    }
                }
            },
            "put": {
                "description": "更新workspace",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "仓库"
                ],
                "summary": "更新workspace",
                "parameters": [
                    {
                        "description": "FileInfo",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.FileInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功后返回值",
                        "schema": {
                            "$ref": "#/definitions/response.FileInfo"
                        }
                    }
                }
            }
        },
        "/user/auth/:repo": {
            "get": {
                "description": "gitee auth login",
                "tags": [
                    "用户"
                ],
                "summary": "授权登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "仓库类型 gitee or github",
                        "name": "repo",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "{redirect_uri}?code=abc",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "description": "获取用户的基本信息",
                "tags": [
                    "用户"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "accesstoken",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.FileInfo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "sha": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "response.FileInfo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "download_url": {
                    "type": "string"
                },
                "html_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "sha": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "url": {
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
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "ShowNote API",
	Description: "This is a sample server Petstore server.",
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
