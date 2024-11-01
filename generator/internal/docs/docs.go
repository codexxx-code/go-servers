// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ilia Ivanov",
            "email": "bonavii@icloud.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/horoscope": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "horoscope"
                ],
                "summary": "Получение гороскопов по знакам зодиака",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "example": [
                            "2024-01-01",
                            "2024-01-02"
                        ],
                        "name": "dates",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "example": [
                            "aries",
                            "taurus",
                            "gemini",
                            "cancer",
                            "leo",
                            "virgo",
                            "libra",
                            "scorpio",
                            "sagittarius",
                            "capricorn",
                            "aquarius",
                            "pisces"
                        ],
                        "name": "zodiacs",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Horoscope"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/horoscope/withGeneration": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "horoscope"
                ],
                "summary": "Получение гороскопа по знакам зодиака, если его нет - генерация",
                "parameters": [
                    {
                        "type": "string",
                        "format": "date",
                        "name": "date",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "aries",
                            "taurus",
                            "gemini",
                            "cancer",
                            "leo",
                            "virgo",
                            "libra",
                            "scorpio",
                            "sagittarius",
                            "capricorn",
                            "aquarius",
                            "pisces"
                        ],
                        "type": "string",
                        "name": "zodiacs",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Horoscope"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/prompt": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prompt"
                ],
                "summary": "Получение промптов по фильтрам",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "cases",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "languages",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Prompt"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prompt"
                ],
                "summary": "Создание промпта по фильтрам",
                "parameters": [
                    {
                        "description": "model.CreatePromptReq",
                        "name": "Body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.CreatePromptReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CreatePromptRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prompt"
                ],
                "summary": "Удаление промпта",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prompt"
                ],
                "summary": "Обновление промпта (patch)",
                "parameters": [
                    {
                        "description": "model.UpdatePromptReq",
                        "name": "Body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePromptReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "Поскольку стандартный энкодер json в го не умеет нормально сериализовать тип ошибок, эта переменная\nИспользуется для подставления значения Err прямо перед сериализацией ошибки в функции JSON",
                    "type": "string"
                },
                "humanText": {
                    "description": "Человекочитаемый текст, который можно показать клиенту\nПеременная настраивается через errors.HumanTextOption(messageWithFmt, args...)\nЕсли значения нет, то автоматически проставляется шаблонными данными в функции middleware.DefaultErrorEncoder",
                    "type": "string"
                },
                "parameters": {
                    "description": "Дополнительные параметры, направленные на дополнение ошибки контекстом, которые проставляются\nЧерез errors.ParamsOption(key1, value1, key2, value2, ...)",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "path": {
                    "description": "Стектрейс от места враппинга ошибки. Если необходимо начать стектрейс с уровня выше, то\nНеобходимо воспользоваться errors.SkipThisCallOption(errors.\u003cconst\u003e)\nconst = SkipThisCall - начать стектрейс на один уровень выше враппера errors.Type.Wrap по дереву\nconst = SkipPreviousCaller и остальные работают по аналогии, пропуская все больше уровней стека вызовов",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "systemInfo": {
                    "description": "Служебное поле, которое автоматически заполняется в функции middleware.DefaultErrorEncoder\nвспомогательными данными из контекста",
                    "$ref": "#/definitions/model.SystemInfo"
                },
                "userInfo": {
                    "description": "Служебное поле, которое автоматически заполняется в функции middleware.DefaultErrorEncoder\nвспомогательными данными из контекста",
                    "$ref": "#/definitions/model.UserInfo"
                }
            }
        },
        "model.CreatePromptReq": {
            "type": "object",
            "required": [
                "case",
                "language",
                "text"
            ],
            "properties": {
                "case": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "model.CreatePromptRes": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.Horoscope": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "format": "date"
                },
                "id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "zodiac": {
                    "type": "string",
                    "enum": [
                        "aries",
                        "taurus",
                        "gemini",
                        "cancer",
                        "leo",
                        "virgo",
                        "libra",
                        "scorpio",
                        "sagittarius",
                        "capricorn",
                        "aquarius",
                        "pisces"
                    ]
                }
            }
        },
        "model.Prompt": {
            "type": "object",
            "properties": {
                "case": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "language": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "model.SystemInfo": {
            "type": "object",
            "properties": {
                "build": {
                    "type": "string"
                },
                "commit": {
                    "type": "string"
                },
                "env": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "model.UpdatePromptReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "case": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "language": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "deviceID": {
                    "type": "string"
                },
                "taskID": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthJWT": {
            "description": "JWT-токен авторизации",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "@{version} (build @{build}) (commit @{commit})",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Horoscope Server Documentation",
	Description:      "API Documentation for Coin",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
