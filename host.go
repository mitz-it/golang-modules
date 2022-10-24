package modules

type Host struct {
	api     *API
	workers []IWorker
}

func (host *Host) Run() {
	var forever chan struct{}

	if host.api != nil {
		go host.api.run()
	}

	for _, worker := range host.workers {
		go worker.Run()
	}

	<-forever
}

func NewHost(api *API, workers []IWorker) *Host {
	return &Host{
		api:     api,
		workers: workers,
	}
}
