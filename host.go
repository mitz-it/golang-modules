package modules

type Host struct {
	api     *API
	modules []*Module
}

func (host *Host) Run() {
	var forever chan struct{}

	for _, module := range host.modules {

		module.invokeInitCalls()

		go module.startWorkers()
	}

	if host.api != nil {
		go host.api.run()
	}
	<-forever
}

func NewHost(api *API, modules []*Module) *Host {
	return &Host{
		api:     api,
		modules: modules,
	}
}
