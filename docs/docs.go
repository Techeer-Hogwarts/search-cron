// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/delete/document": {
            "delete": {
                "description": "Delete intended document in Meilisearch",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delete"
                ],
                "summary": "Delete document in index in meilisearch",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of index",
                        "name": "index",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id of document",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/delete/document/all": {
            "delete": {
                "description": "Delete all documents in an index in Meilisearch. Use at your own risk",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delete"
                ],
                "summary": "Delete all document in meilisearch",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of index",
                        "name": "index",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/delete/index": {
            "delete": {
                "description": "Delete intended index in Meilisearch",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delete"
                ],
                "summary": "Delete index in meilisearch",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of index",
                        "name": "index",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/delete/index/all": {
            "delete": {
                "description": "Delete all index in Meilisearch. Use at your own risk",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delete"
                ],
                "summary": "Delete all index in meilisearch",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/index/blog": {
            "post": {
                "description": "Add Blog Index to Meilisearch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Creates index in meilisearch",
                "parameters": [
                    {
                        "description": "Blog Index",
                        "name": "index",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BlogIndex"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/index/event": {
            "post": {
                "description": "Add Event Index to Meilisearch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Creates index in meilisearch",
                "parameters": [
                    {
                        "description": "Event Index",
                        "name": "index",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EventIndex"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/index/project": {
            "post": {
                "description": "Add Project Index to Meilisearch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Creates index in meilisearch",
                "parameters": [
                    {
                        "description": "Project Index",
                        "name": "index",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProjectIndex"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/index/resume": {
            "post": {
                "description": "Add Resume Index to Meilisearch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Creates index in meilisearch",
                "parameters": [
                    {
                        "description": "Resume Index",
                        "name": "index",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ResumeIndex"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/index/session": {
            "post": {
                "description": "Add Session Index to Meilisearch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Creates index in meilisearch",
                "parameters": [
                    {
                        "description": "Session Index",
                        "name": "index",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SessionIndex"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/index/study": {
            "post": {
                "description": "Add Study Index to Meilisearch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Creates index in meilisearch",
                "parameters": [
                    {
                        "description": "Study Index",
                        "name": "index",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.StudyIndex"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/index/user": {
            "post": {
                "description": "Add User Index to Meilisearch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Creates index in meilisearch",
                "parameters": [
                    {
                        "description": "User Index",
                        "name": "index",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserIndex"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/sync/all": {
            "post": {
                "description": "Sync all index with database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sync"
                ],
                "summary": "Sync all index in meilisearch",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/sync/new": {
            "post": {
                "description": "Sync new index with database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sync"
                ],
                "summary": "Sync new index in meilisearch",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BlogIndex": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2021-01-01"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "stack": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Go"
                    ]
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://example.com/thumbnail.jpg"
                },
                "title": {
                    "type": "string",
                    "example": "블로그 제목"
                },
                "url": {
                    "type": "string",
                    "example": "https://example.com/blog"
                },
                "userID": {
                    "type": "string",
                    "example": "1"
                },
                "userName": {
                    "type": "string",
                    "example": "윤정은"
                },
                "userProfileImage": {
                    "type": "string",
                    "example": "https://example.com/profile.jpg"
                }
            }
        },
        "models.EventIndex": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string",
                    "example": "세션"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "title": {
                    "type": "string",
                    "example": "세션 제목"
                }
            }
        },
        "models.ProjectIndex": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "type": "string",
                    "example": "프로젝트 이름"
                },
                "projectExplain": {
                    "type": "string",
                    "example": "프로젝트 설명"
                },
                "resultImages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "https://example.com/result.jpg"
                    ]
                },
                "teamStacks": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Go"
                    ]
                },
                "title": {
                    "type": "string",
                    "example": "내부용. 무시하세요"
                }
            }
        },
        "models.ResumeIndex": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2021-01-01"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "title": {
                    "type": "string",
                    "example": "이력서 제목"
                },
                "url": {
                    "type": "string",
                    "example": "https://example.com/resume"
                },
                "userID": {
                    "type": "string",
                    "example": "1"
                },
                "userName": {
                    "type": "string",
                    "example": "윤정은"
                },
                "userProfileImage": {
                    "type": "string",
                    "example": "https://example.com/profile.jpg"
                }
            }
        },
        "models.SessionIndex": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2021-01-01"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "likeCount": {
                    "type": "integer",
                    "example": 10
                },
                "presenter": {
                    "type": "string",
                    "example": "윤정은"
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://example.com/thumbnail.jpg"
                },
                "title": {
                    "type": "string",
                    "example": "세션 제목"
                },
                "viewCount": {
                    "type": "integer",
                    "example": 100
                }
            }
        },
        "models.StudyIndex": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "type": "string",
                    "example": "스터디 이름"
                },
                "resultImages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "https://example.com/result.jpg"
                    ]
                },
                "studyExplain": {
                    "type": "string",
                    "example": "스터디 설명"
                },
                "title": {
                    "type": "string",
                    "example": "내부용. 무시하세요"
                }
            }
        },
        "models.UserIndex": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "grade": {
                    "type": "string",
                    "example": "3"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "type": "string",
                    "example": "윤정은"
                },
                "profileImage": {
                    "type": "string",
                    "example": "https://example.com/profile.jpg"
                },
                "school": {
                    "type": "string",
                    "example": "서울대학교"
                },
                "stack": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Go"
                    ]
                },
                "year": {
                    "type": "integer",
                    "example": 7
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
