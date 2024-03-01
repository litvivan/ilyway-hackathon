package app

import (
	"log"
	"time"

	"github.com/litvivan/ilyway/api"
	"github.com/litvivan/ilyway/conf"
	"github.com/litvivan/ilyway/infra/di"

	"github.com/pkg/errors"
)

type Application struct {
	config    *conf.Config
	logger    *log.Logger
	container *di.Container
}

func NewApplication() (*Application, error) {
	logger := log.Default()

	config, err := conf.LoadConfig(".")
	if err != nil {
		return nil, errors.Errorf("failed to load config: %s", err)
	}

	container, err := di.NewContainer(
		&config,
		logger,
		time.Now,
	)
	if err != nil {
		return nil, errors.Errorf("faled to initialize container: %s", err)
	}

	return &Application{
		config:    &config,
		logger:    logger,
		container: container,
	}, nil
}

func (app *Application) Run() error {
	router := api.NewRouter(
		app.config,
		app.container.CatalogService(),
		app.container.CityService(),
		app.logger,
	)

	router.Run()

	return nil
}
