package modules

import (
	"errors"

	"go.uber.org/dig"
)

type ConfigureModule func(config *ModuleConfiguration)

type InitCall func(container *dig.Container)

type ModuleConfiguration struct {
	name            string
	controllersFunc []ControllerConstructorFunc
	workersFunc     []WorkerConstructorFunc
	container       *dig.Container
	initCalls       []InitCall
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

func (config *ModuleConfiguration) SetupInitCall(initCall InitCall) {
	config.initCalls = append(config.initCalls, initCall)
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

func (config *ModuleConfiguration) build() *Module {
	config.validate()
	module := newModule(config.name, config.container)

	for _, controllerFunc := range config.controllersFunc {
		controller := controllerFunc(module.container)
		module.appendControler(controller)
	}

	for _, workerFunc := range config.workersFunc {
		worker := workerFunc(module.container)
		module.appendWorker(worker)
	}

	for _, initCall := range config.initCalls {
		module.appendInitCall(initCall)
	}

	return module
}

func newModuleConfiguration() *ModuleConfiguration {
	config := new(ModuleConfiguration)
	config.controllersFunc = make([]ControllerConstructorFunc, 0)
	config.workersFunc = make([]WorkerConstructorFunc, 0)
	config.initCalls = make([]InitCall, 0)
	return config
}
