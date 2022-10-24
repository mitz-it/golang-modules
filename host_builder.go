package modules

import "errors"

type HostBuilder struct {
	configurations []*ModuleConfiguration
	configureAPI   *ConfigureAPIFunc
}

func (builder *HostBuilder) ConfigureAPI(configure ConfigureAPIFunc) {
	builder.configureAPI = &configure
}

func (builder *HostBuilder) AddModule(configure ConfigureModule) {
	config := NewModuleConfiguration()

	configure(config)

	if config.name == "" {
		err := errors.New("module name cannot be empyt")
		panic(err)
	}

	if config.container == nil {
		err := errors.New("di container cannot be nil")
		panic(err)
	}

	builder.configurations = append(builder.configurations, config)
}

func (builder *HostBuilder) Build() *Host {
	api := new(API)

	if builder.configureAPI != nil {
		api = newAPI()
		configureAPI := *builder.configureAPI
		configureAPI(api)
		api.validate()
		api.buildRootGroup()
	}

	workers := make([]IWorker, 0)

	for _, config := range builder.configurations {
		container := config.container

		if api != nil {
			name := config.name

			group := api.rootGroup.Group(name)

			for _, controllerFunc := range config.controllersFunc {
				controller := controllerFunc(container)
				controller.Register(group)
			}
		}

		for _, workersFunc := range config.workersFunc {
			worker := workersFunc(container)
			workers = append(workers, worker)
		}
	}

	host := NewHost(api, workers)
	return host
}

func NewHostBuilder() *HostBuilder {
	configurations := make([]*ModuleConfiguration, 0)
	return &HostBuilder{
		configureAPI:   nil,
		configurations: configurations,
	}
}
