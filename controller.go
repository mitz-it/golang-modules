package modules

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type ControllerConstructorFunc func(container *dig.Container) IController

type IController interface {
	Register(group *gin.RouterGroup)
}
