package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	catalogsvc "github.com/litvivan/ilyway/api/v1/catalog"
	citiessvc "github.com/litvivan/ilyway/api/v1/cities"
	"github.com/litvivan/ilyway/app/services/catalog"
	"github.com/litvivan/ilyway/app/services/city"
	"github.com/litvivan/ilyway/conf"
)

type Router struct {
	config         *conf.Config
	catalogService *catalog.Service
	cityService    *city.Service
	logger         *log.Logger
}

func NewRouter(
	config *conf.Config,
	catalogService *catalog.Service,
	cityService *city.Service,
	logger *log.Logger,
) *Router {
	return &Router{
		config:         config,
		catalogService: catalogService,
		cityService:    cityService,
		logger:         logger,
	}
}

func (api *Router) Run() {
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	citiesAPISvc := citiessvc.NewService(*api.cityService)
	r.HandleFunc("/cities", citiesAPISvc.GetCities).Methods("GET")

	catalogAPISvc := catalogsvc.NewService(
		*api.cityService,
		*api.catalogService,
	)

	r.HandleFunc("/items", catalogAPISvc.ListItems).Methods("GET", "POST")
	r.HandleFunc("/items/{id:[0-9]+}", catalogAPISvc.GetItem).Methods("GET", "POST")
	r.HandleFunc("/items/seed_service_kek", catalogAPISvc.SeedItems).Methods("GET", "POST")

	http.Handle("/", r)

	listen := fmt.Sprintf(":%d", api.config.WebPort)

	api.logger.Fatal(http.ListenAndServe(listen, nil))
}
