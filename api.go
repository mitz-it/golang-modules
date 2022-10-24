package modules

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type API struct {
	router    *gin.Engine
	rootGroup *gin.RouterGroup
}

func (api *API) useSwaggerHandler() {
	api.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (api *API) Run() {
	api.router.Run()
}

func NewAPI() *API {
	router := gin.Default()
	rootGroup := router.Group("/api")
	return &API{
		router:    router,
		rootGroup: rootGroup,
	}
}
