package di

import (
	"database/sql"
	"log"
	"time"

	cityRepo "github.com/litvivan/ilyway/app/repo/city"
	itemRepo "github.com/litvivan/ilyway/app/repo/item"
	"github.com/litvivan/ilyway/app/services/catalog"
	"github.com/litvivan/ilyway/app/services/city"
	"github.com/litvivan/ilyway/conf"
	"github.com/litvivan/ilyway/infra/db/storage"

	_ "github.com/lib/pq"
)

type Container struct {
	db             *sql.DB
	logger         *log.Logger
	now            func() time.Time
	catalogService *catalog.Service
	cityService    *city.Service
}

func NewContainer(
	config *conf.Config,
	logger *log.Logger,
	now func() time.Time,
) (*Container, error) {
	db, err := sql.Open("postgres", config.DBConnectionString())
	if err != nil {
		panic(err)
	}

	queries := storage.New(db)
	itemRepo := itemRepo.NewRepo(db, queries)
	cityRepo := cityRepo.NewRepo()

	catalogService := catalog.NewService(itemRepo)
	cityService := city.NewService(cityRepo)

	return &Container{
		db:             db,
		logger:         logger,
		now:            now,
		catalogService: catalogService,
		cityService:    cityService,
	}, nil
}

func (c *Container) DB() *sql.DB {
	return c.db
}

func (c *Container) Logger() *log.Logger {
	return c.logger
}

func (c *Container) Now() func() time.Time {
	return c.now
}

func (c *Container) CatalogService() *catalog.Service {
	return c.catalogService
}

func (c *Container) CityService() *city.Service {
	return c.cityService
}
