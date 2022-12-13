package modules

import (
	"errors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type ConfigureAPIFunc func(api *API)

type API struct {
	router             *gin.Engine
	rootGroup          *gin.RouterGroup
	basePath           string
	swaggerSpec        *swag.Spec
	useSwagger         bool
	otelgin_middleware gin.HandlerFunc
}

func (api *API) UseOpenTelemetryMiddleware(serviceName string, opts ...otelgin.Option) {
	api.otelgin_middleware = otelgin.Middleware(serviceName, opts...)
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

	if api.otelgin_middleware != nil {
		api.router.Use(api.otelgin_middleware)
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
