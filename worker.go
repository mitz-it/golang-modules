package modules

import "go.uber.org/dig"

type WorkerConstructorFunc func(container *dig.Container) IWorker

type IWorker interface {
	Run()
}
