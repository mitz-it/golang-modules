package modules

type App struct {
	apiModules    []IApiModule
	workerModules []IWorkerModule
	api           IApi
	basePath      string
}

func (app *App) Start() {
	var forever chan struct{}
	go app.runAPI()
	go app.runWorkers()
	<-forever
}

func (app *App) runAPI() {
	basePath := app.basePath
	app.api.SetSwaggerBasePath(basePath)
	group := app.api.CreateGroup(basePath)

	for _, apiModule := range app.apiModules {
		apiModule.Register(group)
	}

	app.api.ConfigureSwaggerHandler()
	app.api.Run()
}

func (app *App) runWorkers() {
	for _, workerModule := range app.workerModules {
		go workerModule.Start()
	}
}

func newApp(apiModules []IApiModule, workerModules []IWorkerModule, api IApi, basePath string) *App {
	return &App{
		apiModules:    apiModules,
		workerModules: workerModules,
		api:           api,
		basePath:      basePath,
	}
}
