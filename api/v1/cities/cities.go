package cities

import (
	"encoding/json"
	"net/http"

	"github.com/litvivan/ilyway/api/models"
	"github.com/litvivan/ilyway/api/renderer"
	"github.com/litvivan/ilyway/app/services/city"
)

type Service struct {
	cityService city.Service
}

func NewService(cityService city.Service) *Service {
	return &Service{
		cityService: cityService,
	}
}

func (svc *Service) GetCities(w http.ResponseWriter, r *http.Request) {
	cities := svc.cityService.List()

	citiesResp := models.MapCities(cities)

	citiesJSON, err := json.Marshal(citiesResp)
	if err != nil {
		renderer.RenderError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(citiesJSON)
}
