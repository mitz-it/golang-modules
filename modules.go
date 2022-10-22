package modules

import (
	"github.com/gin-gonic/gin"
)

type IApiModule interface {
	Register(group *gin.RouterGroup)
}

type IWorkerModule interface {
	Start()
}
