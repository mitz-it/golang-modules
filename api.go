package modules

import (
	"errors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

type ConfigureAPIFunc func(api *API)

type API struct {
	router      *gin.Engine
	rootGroup   *gin.RouterGroup
	basePath    string
	swaggerSpec *swag.Spec
}

func (api *API) UseSwaggerHandler() {
	api.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (api *API) WithSwaggerSpec(spec *swag.Spec) {
	api.swaggerSpec = spec
}

func (api *API) WithBasePath(basePath string) {
	api.basePath = basePath
}

func (api *API) buildRootGroup() {
	rootGroup := api.router.Group(api.basePath)
	api.rootGroup = rootGroup
}

func (api *API) validate() {
	if api.basePath == "" {
		err := errors.New("base path cannot be empty")
		panic(err)
	}

	if api.swaggerSpec == nil {
		err := errors.New("swag spec cannot be nil")
		panic(err)
	}
}

func (api *API) run() {
	api.router.Run()
}

func newAPI() *API {
	router := gin.Default()
	return &API{
		router:   router,
		basePath: "/api",
	}
}
