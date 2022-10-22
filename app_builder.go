package modules

import (
	"errors"

	"go.uber.org/dig"
)

type APIModuleFunc func(container *dig.Container) IApiModule

type WorkerModuleFunc func(container *dig.Container) IWorkerModule

type apiModules map[*dig.Container]APIModuleFunc

type workerModules map[*dig.Container]WorkerModuleFunc

type AppBuilder struct {
	apiModules    apiModules
	workerModules workerModules
	api           IApi
	apiBasePath   string
}

func (builder *AppBuilder) AppendAPIModule(moduleFunc APIModuleFunc, container *dig.Container) *AppBuilder {
	builder.apiModules[container] = moduleFunc
	return builder
}

func (builder *AppBuilder) AppendWorkerModule(moduleFunc WorkerModuleFunc, container *dig.Container) *AppBuilder {
	builder.workerModules[container] = moduleFunc
	return builder
}

func (builder *AppBuilder) WithAPI(api IApi) *AppBuilder {
	builder.api = api
	return builder
}

func (builder *AppBuilder) WithAPIBasePath(basePath string) *AppBuilder {
	builder.apiBasePath = basePath
	return builder
}

func (builder *AppBuilder) Build() *App {
	if builder.api == nil {
		err := errors.New("api cannot be null")
		panic(err)
	}

	if builder.apiBasePath == "" {
		builder.apiBasePath = "/api"
	}

	apiModules := builder.buildAPIModules()
	workerModules := builder.buildWorkerModules()

	app := newApp(apiModules, workerModules, builder.api, builder.apiBasePath)
	return app
}

func (builder *AppBuilder) buildAPIModules() []IApiModule {
	apiModules := make([]IApiModule, 0)

	for container, apiModuleFunc := range builder.apiModules {

		module := apiModuleFunc(container)
		apiModules = append(apiModules, module)
	}
	return apiModules
}

func (builder *AppBuilder) buildWorkerModules() []IWorkerModule {
	workerModules := make([]IWorkerModule, 0)

	for conainer, workerModuleFunc := range builder.workerModules {
		module := workerModuleFunc(conainer)
		workerModules = append(workerModules, module)
	}

	return workerModules
}

func NewAppBuilder() *AppBuilder {
	apiModules := make(apiModules, 0)
	workerModules := make(workerModules, 0)

	return &AppBuilder{
		apiModules:    apiModules,
		workerModules: workerModules,
		api:           nil,
	}
}
