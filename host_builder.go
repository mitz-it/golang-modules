package modules

type HostBuilder struct {
	configurations []*ModuleConfiguration
	api            *API
}

func (builder *HostBuilder) UseSwaggerHandler() {
	builder.api.useSwaggerHandler()
}

func (builder *HostBuilder) AddModule(configure ConfigureModule) {
	config := NewModuleConfiguration()

	configure(config)

	// if config.name == "" {
	// 	// TODO: panic, name cannot be empty
	// }

	// if config.container == nil {
	// 	// TODO: panic, container cannot be nil
	// }

	builder.configurations = append(builder.configurations, config)
}

func (builder *HostBuilder) Build() *Host {
	workers := make([]IWorker, 0)

	for _, config := range builder.configurations {
		container := config.container
		name := config.name

		group := builder.api.rootGroup.Group(name)

		for _, controllerFunc := range config.controllersFunc {
			controller := controllerFunc(container)
			controller.Register(group)
		}

		for _, workersFunc := range config.workersFunc {
			worker := workersFunc(container)
			workers = append(workers, worker)
		}
	}

	host := NewHost(builder.api, workers)
	return host
}

func NewHostBuilder() *HostBuilder {
	api := NewAPI()
	configurations := make([]*ModuleConfiguration, 0)
	return &HostBuilder{
		api:            api,
		configurations: configurations,
	}
}
