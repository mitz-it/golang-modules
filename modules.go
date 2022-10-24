package modules

import (
	"errors"

	"go.uber.org/dig"
)

type ConfigureModule func(config *ModuleConfiguration)

type ModuleConfiguration struct {
	name            string
	controllersFunc []ControllerConstructorFunc
	workersFunc     []WorkerConstructorFunc
	container       *dig.Container
}

func (config *ModuleConfiguration) WithName(name string) {
	config.name = name
}

func (config *ModuleConfiguration) AddController(controllerFunc ControllerConstructorFunc) {
	config.controllersFunc = append(config.controllersFunc, controllerFunc)
}

func (config *ModuleConfiguration) AddWorker(workerFunc WorkerConstructorFunc) {
	config.workersFunc = append(config.workersFunc, workerFunc)
}

func (config *ModuleConfiguration) WithDIContainer(container *dig.Container) {
	config.container = container
}

func (config *ModuleConfiguration) validate() {
	if config.name == "" {
		err := errors.New("module name cannot be empyt")
		panic(err)
	}

	if config.container == nil {
		err := errors.New("di container cannot be nil")
		panic(err)
	}
}

func (config *ModuleConfiguration) registerControllers(api *API) {
	name := config.name
	container := config.container
	group := api.rootGroup.Group(name)

	for _, controllerFunc := range config.controllersFunc {
		controller := controllerFunc(container)
		controller.Register(group)
	}
}

func (config *ModuleConfiguration) createWorkers() []IWorker {
	workers := newWorkers()
	container := config.container
	for _, workersFunc := range config.workersFunc {
		worker := workersFunc(container)
		workers = append(workers, worker)
	}
	return workers
}

func newModuleConfiguration() *ModuleConfiguration {
	config := new(ModuleConfiguration)
	config.controllersFunc = make([]ControllerConstructorFunc, 0)
	config.workersFunc = make([]WorkerConstructorFunc, 0)
	return config
}
