package app

import (
	"context"

	"github.com/solumD/go-service-template/internal/closer"
	"github.com/solumD/go-service-template/internal/config"
)

const configPath = ".env"

// App object of an app
type App struct {
	serviceProvider *serviceProvider
	// servers or handlers
}

// NewApp returns new App object
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run starts an App
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	// some gorutines with running servers
	return nil
}

func (a *App) initDeps(_ context.Context) error {
	err := a.initConfig()
	if err != nil {
		return err
	}

	a.initServiceProvider()
	// some inits

	return nil
}

func (a *App) initConfig() error {
	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider() {
	a.serviceProvider = NewServiceProvider()
}
