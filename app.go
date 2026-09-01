package core

import (
	"sync"
)

type Application struct {
	services map[string]Service
}

func (a *Application) Run() {
	var wg sync.WaitGroup
	defer wg.Done()
	for _, s := range a.services {
		wg.Go(s.Run)
	}
	wg.Wait()
}

func (a *Application) UseServices(builder func(services map[string]Service)) {
	builder(a.services)
}

func NewApplication() *Application {
	return &Application{
		services: make(map[string]Service),
	}
}
