package app

import (
	"sync"

	"github.com/deeerain/nebula-core/service"
)

type Application struct {
	services map[string]service.Service
}

func (a *Application) Run() {
	var wg sync.WaitGroup
	defer wg.Done()
	for _, s := range a.services {
		wg.Go(s.Run)
	}
	wg.Wait()
}

func (a *Application) UseServices(builder func(services map[string]service.Service)) {
	builder(a.services)
}

func NewApplication() *Application {
	return &Application{
		services: make(map[string]service.Service),
	}
}
