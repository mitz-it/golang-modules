package modules

type HostBuilder struct {
	configurations []*ModuleConfiguration
	configureAPI   *ConfigureAPIFunc
	useAPI         bool
}

func (builder *HostBuilder) UseAPI(useAPI bool) {
	builder.useAPI = useAPI
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
	api := builder.buildAPI()

	modules := make([]*Module, 0)

	for _, config := range builder.configurations {
		module := config.build()

		if api != nil {
			module.registerControllers(api)
		}

		modules = append(modules, module)
	}

	host := NewHost(api, modules)
	return host
}

func (builder *HostBuilder) buildAPI() *API {
	if !builder.useAPI {
		return nil
	}

	api := NewAPI()

	if builder.configureAPI != nil {
		configureAPI := *builder.configureAPI
		configureAPI(api)
	}

	api.validate()
	api.configure()

	return api
}

func NewHostBuilder() *HostBuilder {
	configurations := make([]*ModuleConfiguration, 0)
	return &HostBuilder{
		configureAPI:   nil,
		configurations: configurations,
		useAPI:         true,
	}
}

func NewHostBuilder2() *HostBuilder {
	configurations := make([]*ModuleConfiguration, 0)
	return &HostBuilder{
		configureAPI:   nil,
		configurations: configurations,
		useAPI:         true,
	}
}
