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
	useSwagger  bool
}

func (api *API) UseSwagger(spec *swag.Spec) {
	api.useSwagger = true
	api.swaggerSpec = spec
}

func (api *API) WithBasePath(basePath string) {
	api.basePath = basePath
}

func (api *API) validate() {
	if api.useSwagger {
		if api.swaggerSpec == nil {
			err := errors.New("swag spec cannot be nil")
			panic(err)
		}
	}
}

func (api *API) configure() {
	basePath := api.basePath
	rootGroup := api.router.Group(basePath)
	api.rootGroup = rootGroup

	if api.useSwagger {
		api.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.swaggerSpec.BasePath = basePath
	}
}

func (api *API) run() {
	api.router.Run()
}

func NewAPI() *API {
	router := gin.Default()
	return &API{
		router:     router,
		basePath:   "/api",
		useSwagger: false,
	}
}
