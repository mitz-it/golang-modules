package modules

import "github.com/gin-gonic/gin"

type IApi interface {
	SetSwaggerBasePath(basePath string)
	CreateGroup(group string) *gin.RouterGroup
	ConfigureSwaggerHandler()
	Run()
}
