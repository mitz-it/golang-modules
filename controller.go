package modules

import "github.com/gin-gonic/gin"

type IController interface {
	Register(group *gin.RouterGroup)
}
