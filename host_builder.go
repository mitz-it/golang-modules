package modules

type HostBuilder struct {
	configurations []*ModuleConfiguration
	configureAPI   *ConfigureAPIFunc
}

func (builder *HostBuilder) ConfigureAPI(configure ConfigureAPIFunc) {
	builder.configureAPI = &configure
}

func (builder *HostBuilder) AddModule(configure ConfigureModule) {
	config := newModuleConfiguration()

	configure(config)

	builder.configurations = append(builder.configurations, config)
}

func (builder *HostBuilder) Build() *Host {
	api := builder.createAPI()

	workers := newWorkers()

	for _, config := range builder.configurations {
		config.validate()

		if api != nil {
			config.registerControllers(api)
		}

		moduleWorkers := config.createWorkers()
		workers = append(workers, moduleWorkers...)
	}

	host := NewHost(api, workers)
	return host
}

func (builder *HostBuilder) createAPI() *API {
	if builder.configureAPI == nil {
		return nil
	}

	api := NewAPI()

	configureAPI := *builder.configureAPI
	configureAPI(api)

	api.validate()
	api.configure()

	return api
}

func NewHostBuilder() *HostBuilder {
	configurations := make([]*ModuleConfiguration, 0)
	return &HostBuilder{
		configureAPI:   nil,
		configurations: configurations,
	}
}
